package mangosock

import (
	"time"

	"nanomsg.org/go-mangos"
	"nanomsg.org/go-mangos/protocol/rep"
	"nanomsg.org/go-mangos/protocol/req"
	"github.com/redsift/go-mangosock/nano"
)

// ensure we implement interfaces correctly
var (
	_ nano.Req = &reqsock{}
)

type reqsock struct {
	s
}

func NewReqSocket() (nano.Req, error) {
	sock, err := req.NewSocket()
	if err != nil {
		return nil, err
	}

	err = sock.SetOption(mangos.OptionMaxRecvSize, 0) // remove max buffer size limit for recv
	if err != nil {
		return nil, err
	}

	return &reqsock{s: s{sock: sock}}, nil
}

func (s *reqsock) SetResendInterval(interval time.Duration) error {
	return s.sock.SetOption(mangos.OptionRetryTime, interval)
}

// ensure we implement interfaces correctly
var (
	_ nano.Rep = &repsock{}
)

type repsock struct {
	s
}

func NewRepSocket() (nano.Rep, error) {
	sock, err := rep.NewSocket()
	if err != nil {
		return nil, err
	}

	err = sock.SetOption(mangos.OptionMaxRecvSize, 0) // remove max buffer size limit for recv
	if err != nil {
		return nil, err
	}

	return &repsock{s: s{sock: sock}}, nil
}

func (s *repsock) Address() string {
	return s.s.addr
}
