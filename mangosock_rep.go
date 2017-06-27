//
// hilt/mangosock/mangosock_rep.go
//
//
// Copyright (c) 2016 Redsift Limited. All rights reserved.
//

package mangosock

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/rep"
	"github.com/redsift/go-socket"
)

// ensure we implement interfaces correctly
var (
	_ socket.Socket = &MangoRepSock{}
)

type MangoRepSock struct {
	MangoSock
}

func NewRepSocket() (socket.Socket, error) {
	sock, err := rep.NewSocket()
	if err != nil {
		return nil, err
	}

	err = sock.SetOption(mangos.OptionMaxRecvSize, 0) // remove max buffer size limit for recv
	if err != nil {
		return nil, err
	}

	return &MangoRepSock{MangoSock: MangoSock{sock: sock}}, nil
}
