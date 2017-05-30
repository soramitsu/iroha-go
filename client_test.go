package iroha

import (
	"fmt"
	"net"
	"testing"

	"github.com/soramitsu/iroha-go/mock"
	"github.com/soramitsu/iroha-go/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	host := "localhost"
	port := "50051"
	addr := fmt.Sprintf("%s:%s", host, port)
	l, err := net.Listen("tcp", addr)

	require.NoError(err)
	srv := mock.NewMockServer()
	go func() {
		require.NoError(srv.Serve(l))
	}()

	client := NewClient(host, port, nil, grpc.WithInsecure())

	ctx := context.Background()
	a := model.Account{}
	assert.NoError(client.AddAccount(ctx, a))
}
