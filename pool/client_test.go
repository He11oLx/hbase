package pool

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/He11oLx/hbase"
	"log"
)

const (
	host       = "localhost"
	port       = "9090"
	initialCap = 3
	maxCap     = 10
)

func ExampleNewTPoolClient() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	poolClient, err := NewTPoolClient(host, port, protocolFactory, protocolFactory, initialCap, maxCap)
	if err != nil {
		log.Fatalln(err)
	}
	defer poolClient.Destroy()
	client := hbase.NewClient(poolClient)
	tables, err := client.GetTableNames(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(tables) > 0)
	// Output:
	// true
}
