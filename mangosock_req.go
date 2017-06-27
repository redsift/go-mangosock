//
// hilt/mangosock/mangosock_req.go
//
//
// Copyright (c) 2016 Redsift Limited. All rights reserved.
//

package mangosock

import (
	"time"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/req"
	"github.com/redsift/go-socket"
)

// ensure we implement interfaces correctly
var (
	_ socket.Socket = &MangoReqSock{}
)

type MangoReqSock struct {
	MangoSock
}

func NewReqSocket() (socket.Socket, error) {
	sock, err := req.NewSocket()
	if err != nil {
		return nil, err
	}

	err = sock.SetOption(mangos.OptionMaxRecvSize, 0) // remove max buffer size limit for recv
	if err != nil {
		return nil, err
	}

	return &MangoReqSock{MangoSock: MangoSock{sock: sock}}, nil
}

func (s *MangoReqSock) SetResendInterval(interval time.Duration) error {
	return s.sock.SetOption(mangos.OptionRetryTime, interval)
}
