package iroha

import (
	"log"

	"fmt"

	"github.com/soramitsu/iroha-go/iroha"
	"google.golang.org/grpc"
)

type Client struct {
	Endpoint string

	Options []grpc.DialOption
	Logger  *log.Logger
}

func NewClient(host string, port int, options []grpc.DialOption, logger *log.Logger) *Client {
	endpoint := fmt.Sprintf("%s:%d", host, port)

	if logger == nil {
		lgr := log.Logger{}
		logger = &lgr
	}

	return &Client{
		Endpoint: endpoint,
		Options:  options,
		Logger:   logger,
	}
}

func (c *Client) getClient() (iroha.SumeragiClient, error) {
	conn, err := grpc.Dial(c.Endpoint, c.Options...)
	if err != nil {
		return nil, err
	}

	return iroha.NewSumeragiClient(conn), nil
}

