package mangosock

import (
	"time"

	"github.com/redsift/go-mangosock/nano"
	"go.nanomsg.org/mangos/v3"
	"go.nanomsg.org/mangos/v3/transport"
	"go.nanomsg.org/mangos/v3/transport/ipc"
	"go.nanomsg.org/mangos/v3/transport/tcp"
)

var _ nano.Socket = &s{}

type s struct {
	sock mangos.Socket
	addr string
}

func init() {
	transport.RegisterTransport(ipc.Transport)
	transport.RegisterTransport(tcp.Transport)
}

func (s *s) Bind(addr string) error {
	s.addr = addr
	return s.sock.Listen(addr)
}

func (s *s) Connect(addr string) error {
	s.addr = addr
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
