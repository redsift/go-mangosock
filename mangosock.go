//
// hilt/mangosock/mangosock.go
//
//
// Copyright (c) 2016 Redsift Limited. All rights reserved.
//

package mangosock

import (
	"strings"
	"sync"
	"time"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/redsift/go-socket"
)

// ensure we implement interfaces correctly
var (
	_ socket.Socket = &MangoSock{}
)

type MangoSock struct {
	sync.Mutex
	sock mangos.Socket
}

func (s *MangoSock) Bind(addr string) error {
	if strings.HasPrefix(addr, "ipc://") {
		s.sock.AddTransport(ipc.NewTransport())
	}
	return s.sock.Listen(addr)
}

func (s *MangoSock) Connect(addr string) error {
	if strings.HasPrefix(addr, "ipc://") {
		s.sock.AddTransport(ipc.NewTransport())
	}
	return s.sock.Dial(addr)
}

func (s *MangoSock) SetSendTimeout(timeout time.Duration) error {
	return s.sock.SetOption(mangos.OptionSendDeadline, timeout)
}

func (s *MangoSock) SetRecvTimeout(timeout time.Duration) error {
	return s.sock.SetOption(mangos.OptionRecvDeadline, timeout)
}

func (s *MangoSock) SetResendInterval(timeout time.Duration) error {
	return nil
}

func (s *MangoSock) Send(data []byte) error {
	return s.sock.Send(data)
}

func (s *MangoSock) Recv() ([]byte, error) {
	return s.sock.Recv()
}

func (s *MangoSock) Close() error {
	return s.sock.Close()
}
