package iroha

import (
	"log"

	"fmt"

	"github.com/soramitsu/iroha-go/iroha"
	"google.golang.org/grpc"
	"github.com/google/flatbuffers/go"
)

type Client struct {
	Endpoint string

	Options []grpc.DialOption
	Logger  *log.Logger
}

func NewClient(host string, port string, logger *log.Logger, options ...grpc.DialOption) *Client {
	endpoint := fmt.Sprintf("%s:%s", host, port)

	if logger == nil {
		lgr := log.Logger{}
		logger = &lgr
	}

	options = append(options, grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))

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

