package mangosock

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pub"
	"github.com/go-mangos/mangos/protocol/sub"
	"github.com/redsift/go-mangosock/nano"
)

// ensure we implement interfaces correctly
var (
	_ nano.Pub = &pubsock{}
)

type pubsock struct {
	s
}

func NewPubSocket() (nano.Pub, error) {
	sock, err := pub.NewSocket()
	if err != nil {
		return nil, err
	}

	err = sock.SetOption(mangos.OptionMaxRecvSize, 0) // remove max buffer size limit for recv
	if err != nil {
		return nil, err
	}

	return &pubsock{s: s{sock: sock}}, nil
}

func (s *pubsock) Publish(data []byte) (int, error) {
	return s.Send(data)
}

// ensure we implement interfaces correctly
var (
	_ nano.Sub = &subsock{}
)

type subsock struct {
	s
}

func NewSubSocket() (nano.Sub, error) {
	sock, err := sub.NewSocket()
	if err != nil {
		return nil, err
	}

	err = sock.SetOption(mangos.OptionMaxRecvSize, 0) // remove max buffer size limit for recv
	if err != nil {
		return nil, err
	}

	return &subsock{s: s{sock: sock}}, nil
}

func (s *subsock) Subscribe(topic []byte) error {
	return s.sock.SetOption(mangos.OptionSubscribe, topic)
}
