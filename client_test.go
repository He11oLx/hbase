package hbase

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"testing"
)

var host = "localhost"
var port = "9090"
var protocol = "binary"

func TestNewClient(t *testing.T) {
	trans, err := thrift.NewTSocket(net.JoinHostPort(host, port))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		t.Error("Invalid protocol specified: ", protocol)
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := NewClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		t.Error("Error opening socket to ", host, ":", port, " ", err)
	}
	ctx, _ := context.WithCancel(context.Background())

	tbs, _ := client.GetTableNames(ctx)
	for _, tb := range tbs {
		fmt.Println(string(tb))
	}
}
