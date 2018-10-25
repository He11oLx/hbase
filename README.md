# HBase
Use HBase1.20 in Golang1.11 by thrift1.

thrift package install：
- https://github.com/apache/thrift/tree/0.11.0/lib/go
- go get github.com/apache/thrift/lib/go/thrift
- cd $GOPATH/src/git.apache.org/thrift.git && git checkout 0.11.0

HBase DSL：
- DSL download：https://github.com/apache/hbase/tree/master/hbase-thrift/src/main/resources/org/apache/hadoop/hbase
- compiler download：http://thrift.apache.org/download
- Careful! The generated hbase.go file does not apply to Go 1.11.
- Maybe you should modify it yourself or use mine directly.

```
Drop some word which has prefix "hbase".

Replace:
\{\s+temp\s+:=(.+)\s+([^=]+)=\s+temp\s+}    ->  {\n\t$2 = $1\t\n}

Text        --- replace to (Match words && case)-->     []byte
Bytes       --- replace to (Match words && case)-->     []byte
[]byte(v)   --- replace to (Match words && case)-->     v

Drop:
func.*(GetColumns|GetValues|GetAttributes|GetRow|GetTableName)\(\).*\s+.*\s+\}
func.*(GetTable|GetRows|GetColumn|GetTimestamp|GetNumVersions)\(\).*\s+.*\s+\}
func.*(GetMutations|GetValue)\(\).*\s+.*\s+\}

Replace:
!p.IsSet(\w+)\(\)   ->  p.$1 == nil
p.IsSet(\w+)\(\)    ->  p.$1 != nil

Drop：
func.*Is\w+\(\).*\s+.*\s+\}
```

## StandardClient-Demo
```
package main

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/He11oLx/hbase"
	"log"
	"net"
)

const (
	host = "localhost"
	port = "9090"
)

func main() {
	trans, _ := thrift.NewTSocket(net.JoinHostPort(host, port))
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := hbase.NewClient(thrift.NewTStandardClient(iprot, oprot))
	trans.Open()
	defer trans.Close()

	tables, err := client.GetTableNames(context.Background())
	if err != nil {
		log.Fatalln(err)
	} else {
		for _, t := range tables {
			log.Println(string(t))
		}
	}
}

```

## PoolClient-Demo
```
package main

import (
	"context"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/He11oLx/hbase"
	"github.com/He11oLx/hbase/pool"
	"log"
)

const (
	host       = "localhost"
	port       = "9090"
	initialCap = 3
	maxCap     = 10
)

func main() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	poolClient, err := pool.NewTPoolClient(host, port, protocolFactory, protocolFactory, initialCap, maxCap)
	if err != nil {
		log.Fatalln(err)
	}
	defer poolClient.Destroy()

	client := hbase.NewClient(poolClient)
	tables, err := client.GetTableNames(context.Background())
	if err != nil {
		log.Fatalln(err)
	} else {
		for _, t := range tables {
			log.Println(string(t))
		}
	}
}

```