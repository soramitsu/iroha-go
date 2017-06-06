package iroha

import (
	"log"
	"os"

	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/protocol"
	"google.golang.org/grpc"
)

type Client struct {
	Target  string
	Logger  *log.Logger
	Options []grpc.DialOption
}

func NewClient(target string, logger *log.Logger, options ...grpc.DialOption) *Client {
	if logger == nil {
		logger = log.New(os.Stdout, "iroha-go", log.LstdFlags)
	}

	options = append(options, grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))

	return &Client{
		Target:  target,
		Options: options,
		Logger:  logger,
	}
}

func (c *Client) getClient() (protocol.ClientServiceClient, error) {
	conn, err := grpc.Dial(c.Target, c.Options...)
	if err != nil {
		return nil, err
	}

	return protocol.NewClientServiceClient(conn), nil
}
