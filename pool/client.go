package pool

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"time"
)

var (
	DefaultConnectTimeout = time.Duration(time.Second * 30)
	DefaultIdleTimeout    = time.Duration(time.Hour * 1)
	MaxRetryGet           = 2
)

type TPoolClient struct {
	*thrift.TStandardClient
	// we can not use thrift.client.seqId
	seqId                      int32
	timeout                    time.Duration
	iprotFactory, oprotFactory thrift.TProtocolFactory
	pool                       Pool
}

func NewTPoolClient(host, port string, inputProtocol, outputProtocol thrift.TProtocolFactory, initialCap, maxCap int) (*TPoolClient, error) {
	newFunc := func() (interface{}, error) {
		conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			return nil, err
		}
		return conn, err
	}
	closeFunc := func(v interface{}) error { return v.(net.Conn).Close() }

	p, err := NewChannelPool(initialCap, maxCap, newFunc, closeFunc, nil, DefaultIdleTimeout)
	if err != nil {
		return nil, err
	}
	return &TPoolClient{
		iprotFactory: inputProtocol,
		oprotFactory: outputProtocol,
		pool:         p,
		timeout:      DefaultConnectTimeout,
	}, nil
}

func (p *TPoolClient) SetTimeout(timeout time.Duration) error {
	p.timeout = timeout
	return nil
}

func (p *TPoolClient) Call(ctx context.Context, method string, args, result thrift.TStruct) error {
	var err error
	var connVar interface{}
	for i := 0; i < MaxRetryGet; i++ {
		connVar, err = p.pool.Get()
		if err == nil {
			break
		} else if err == ErrDestroyed {
			return err
		}
	}
	if connVar == nil {
		connVar, err = p.pool.New()
		if err != nil {
			return err
		}
	}

	p.seqId++
	seqId := p.seqId
	conn := connVar.(net.Conn)

	trans := thrift.NewTSocketFromConnTimeout(conn, p.timeout)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	inputProtocol := protocolFactory.GetProtocol(trans)
	outputProtocol := protocolFactory.GetProtocol(trans)

	if err = p.Send(outputProtocol, seqId, method, args); err != nil {
		return err
	}
	if result == nil {
		p.pool.Put(connVar)
		return nil
	}

	if err = p.Recv(inputProtocol, seqId, method, result); err != nil {
		return err
	}
	p.pool.Put(connVar)
	return nil
}

func (p *TPoolClient) Destroy() {
	p.pool.Destroy()
}
