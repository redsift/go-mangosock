package nano

import "time"

type Socket interface {
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
	Socket
	Address() string
}

type Req interface {
	Socket
	SetResendInterval(timeout time.Duration) error
}

type Sub interface {
	Socket
	Subscribe(string) error
	Unsubscribe(string) error
}

type Pub interface {
	Socket
	Publish([]byte) (int, error)
}
