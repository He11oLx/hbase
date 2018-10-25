package pool

import (
	"fmt"
	"testing"
	"time"
)

func TestNewChannelPool(t *testing.T) {
	seq := 0
	pool, e := NewChannelPool(1, 3,
		func() (interface{}, error) {
			seq++
			return seq, nil
		},
		func(i interface{}) error {
			fmt.Printf("close %d\n", i)
			return nil
		},
		nil,
		time.Duration(time.Second*5),
	)
	if e != nil {
		t.Fatal(e)
	}
	defer pool.Destroy()

	fmt.Printf("Init channel length:%d\n", pool.Len())
	i1, _ := pool.Get()
	if i1 != 1 {
		t.Fatalf("wanted 1,go %d", i1)
	}
	fmt.Printf("After Got() length:%d\n", pool.Len())

	i2, _ := pool.Get()
	if i2 != 2 {
		t.Fatalf("wanted 2,go %d", i1)
	}
	fmt.Printf("After ReGot() length:%d\n", pool.Len())
	pool.Put(i1)
	pool.Put(i2)
	pool.Put(33)
	pool.Put(44)
	fmt.Printf("After Put() length:%d\n", pool.Len())
}

func TestChannelPool_TimeOut(t *testing.T) {
	seq := 0
	pool, _ := NewChannelPool(1, 1,
		func() (interface{}, error) {
			seq++
			return seq, nil
		},
		func(i interface{}) error {
			fmt.Printf("timeout close conn：%d\n", i)
			return nil
		},
		nil,
		time.Duration(time.Second*5),
	)
	time.Sleep(time.Second * 6)
	i1, err := pool.Get()
	if err == ErrTimeOut {
		fmt.Printf("After TimeOut Got()=> %s\n", err)
	} else {
		fmt.Printf("After TimeOut Got()=> %d\n", i1)
	}
}

func TestChannelPool_Ping(t *testing.T) {
	seq := 0
	pool, _ := NewChannelPool(3, 3,
		func() (interface{}, error) {
			seq++
			return seq, nil
		},
		func(i interface{}) error {
			fmt.Printf("ping close conn：%d\n", i)
			return nil
		},
		func(i interface{}) error {
			if i == 3 {
				fmt.Printf("ErrPing close conn：%d\n", i)
				return ErrPing
			}
			return nil
		},
		time.Duration(time.Hour*10),
	)
	pool.Get()
	pool.Get()
	i3, err := pool.Get()
	if err == ErrPing {
		fmt.Printf("After ErrPing Got()=> %s\n", err)
	} else {
		fmt.Printf("After ErrPing Got()=> %d\n", i3)
	}
}
