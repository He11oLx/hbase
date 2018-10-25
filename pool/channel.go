package pool

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var pools []*channelPool

type New func() (interface{}, error)
type Close func(interface{}) error
type Ping func(interface{}) error

type channelPool struct {
	mu      sync.Mutex
	conns   chan *idleConn
	new     New
	close   Close
	ping    Ping
	timeout time.Duration
}

type idleConn struct {
	conn interface{}
	t    time.Time
}

// new,close 不能为 nil, ping 可以为 nil
func NewChannelPool(initCap, maxCap int, new New, close Close, ping Ping, timeout time.Duration) (Pool, error) {
	if initCap < 0 || maxCap <= 0 || initCap > maxCap {
		return nil, errors.New("invalid capacity settings")
	}
	if new == nil {
		return nil, errors.New("invalid factory func settings")
	}
	if close == nil {
		return nil, errors.New("invalid close func settings")
	}

	c := &channelPool{
		conns:   make(chan *idleConn, maxCap),
		new:     new,
		close:   close,
		timeout: timeout,
		ping:    ping,
	}

	for i := 0; i < initCap; i++ {
		conn, err := c.new()
		if err != nil {
			c.Destroy()
			return nil, fmt.Errorf("can not create new conn: %s", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}

	pools = append(pools, c)
	return c, nil
}

func (c *channelPool) New() (interface{}, error) {
	conn, err := c.new()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *channelPool) Get() (interface{}, error) {
	conns := c.conns
	if conns == nil {
		return nil, ErrDestroyed
	}
	select {
	case wrapConn := <-conns:
		// 判断是否超时，超时则丢弃
		if timeout := c.timeout; timeout > 0 {
			if wrapConn.t.Add(timeout).Before(time.Now()) {
				// 丢弃并关闭该连接，忽略了实际错误
				c.close(wrapConn.conn)
				return nil, ErrTimeOut
			}
		}
		// 判断是否失效，失效则丢弃，如果用户没有设定 ping 方法，就不检查
		if c.ping != nil {
			if err := c.Ping(wrapConn.conn); err != nil {
				// 忽略实际错误
				return nil, ErrPing
			}
		}
		return wrapConn.conn, nil
	default:
		return c.new()
	}
}

func (c *channelPool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conns == nil {
		return c.close(conn)
	}

	select {
	case c.conns <- &idleConn{conn: conn, t: time.Now()}:
		return nil
	default:
		// 连接池已满，直接关闭该连接
		return c.close(conn)
	}
}

func (c *channelPool) Ping(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	if c.ping == nil {
		return errors.New("ping func is nil. rejecting")
	}
	return c.ping(conn)
}

func (c *channelPool) Destroy() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.new = nil
	closeFun := c.close
	c.close = nil
	c.ping = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}

	close(conns)
	for wrapConn := range conns {
		closeFun(wrapConn.conn)
	}
}

func (c *channelPool) Len() int {
	return len(c.conns)
}

// 摧毁所有已创建的 pool
func DestroyAllPool() {
	for _, p := range pools {
		p.Destroy()
	}
}
