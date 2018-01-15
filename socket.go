package mangosock

import (
	"strings"
	"time"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"github.com/redsift/go-mangosock/nano"
)

var _ nano.Socket = &s{}

type s struct {
	sock mangos.Socket
	addr string
}

func (s *s) addTransport(addr string) {
	if strings.HasPrefix(addr, "ipc://") {
		s.sock.AddTransport(ipc.NewTransport())
	}
	if strings.HasPrefix(addr, "tcp://") {
		s.sock.AddTransport(tcp.NewTransport())
	}
}

func (s *s) Bind(addr string) error {
	s.addr = addr
	s.addTransport(addr)
	return s.sock.Listen(addr)
}

func (s *s) Connect(addr string) error {
	s.addr = addr
	s.addTransport(addr)
	return s.sock.Dial(addr)
}

func (s *s) SetSendTimeout(timeout time.Duration) error {
	return s.sock.SetOption(mangos.OptionSendDeadline, timeout)
}

func (s *s) SetRecvTimeout(timeout time.Duration) error {
	return s.sock.SetOption(mangos.OptionRecvDeadline, timeout)
}

func (s *s) SetRecvMaxSize(size int64) error {
	sz := int(size)
	if size < 0 {
		sz = 0
	}

	return s.sock.SetOption(mangos.OptionMaxRecvSize, sz)
}

func (s *s) Send(data []byte) (int, error) {
	err := s.sock.Send(data)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func (s *s) Recv() ([]byte, error) {
	return s.sock.Recv()
}

func (s *s) Close() error {
	return s.sock.Close()
}
