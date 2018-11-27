package pool

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"time"
)

var (
	DefaultConnectTimeout = time.Duration(time.Second * 1)
	DefaultIdleTimeout    = time.Duration(time.Hour * 1)
	DefaultMaxRetry       = 2
)

type TPoolClient struct {
	*thrift.TStandardClient
	// we can not use thrift.client.seqId
	seqId                      int32
	timeout                    time.Duration
	iprotFactory, oprotFactory thrift.TProtocolFactory
	pool                       Pool
	maxRetry                   int
}

func NewTPoolClient(host, port string, inputProtocol, outputProtocol thrift.TProtocolFactory, initialCap, maxCap int) (*TPoolClient, error) {
	newFunc := func() (interface{}, error) {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), DefaultConnectTimeout)
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
		maxRetry:     DefaultMaxRetry,
	}, nil
}

func (p *TPoolClient) SetTimeout(timeout time.Duration) {
	p.timeout = timeout
}
func (p *TPoolClient) SetMaxRetry(maxRetry int) {
	if maxRetry > 1 {
		p.maxRetry = maxRetry
	}
}

func (p *TPoolClient) Call(ctx context.Context, method string, args, result thrift.TStruct) (err error) {
	var connVar interface{}
	// try old conn
	for i := 0; i < p.maxRetry; i++ {
		connVar, err = p.pool.Get()
		// 闲置时间过长，会自动断开 conVar 为 nil
		if connVar == nil {
			continue
		}
		if err == nil {
			err = p.call(connVar, method, args, result)
		}
		switch err.(type) {
		case nil:
			p.pool.Put(connVar)
			return nil
		case thrift.TTransportException:
			// TTransportException 是 TProtocolException 的子集
			p.pool.Close(connVar.(net.Conn))
			return err
		case thrift.TProtocolException:
			// TProtocolException 是 TException 的子集
			p.pool.Close(connVar.(net.Conn))
			continue
		default:
			p.pool.Close(connVar.(net.Conn))
			return err
		}
	}

	// try new conn
	connVar, err = p.pool.New()
	if err != nil {
		return err
	}
	return p.call(connVar, method, args, result)
}
func (p *TPoolClient) call(connVar interface{}, method string, args, result thrift.TStruct) (err error) {
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
		return
	}

	if err = p.Recv(inputProtocol, seqId, method, result); err != nil {
		return err
	}
	return
}

func (p *TPoolClient) Destroy() {
	p.pool.Destroy()
}
