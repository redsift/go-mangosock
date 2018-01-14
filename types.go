package mangosock

import "time"

type socket interface {
	Bind(addr string) error
	Connect(addr string) error
	SetSendTimeout(timeout time.Duration) error
	SetRecvTimeout(timeout time.Duration) error
	SetRecvMaxSize(int64) error
	Send(data []byte) (int, error)
	Recv() ([]byte, error)
	Close() error
}

type Rep interface {
	socket
	Address() string
}

type Req interface {
	socket
	SetResendInterval(timeout time.Duration) error
}

type Sub interface {
	socket
	Subscribe([]byte) error
}

type Pub interface {
	socket
	Publish([]byte) (int, error)
}
