package pool

import "errors"

var (
	ErrDestroyed = errors.New("pool is closed")
	ErrTimeOut = errors.New("conn timeout")
	ErrPing = errors.New("conn ping error")
)

type Pool interface {

	New() (interface{}, error)

	Get() (interface{}, error)

	Put(interface{}) error

	Destroy()

	Len() int
}
