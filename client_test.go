package iroha

import (
	"net"
	"testing"

	"github.com/soramitsu/iroha-go/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {

	target := "localhost:50051"
	l, err := net.Listen("tcp", target)
	require.NoError(t, err)

	srv := mock.NewMockServer()
	go func() {
		require.NoError(t, srv.Serve(l))
	}()

	client := NewClient(target, nil, grpc.WithInsecure())

	_, err = client.getClient()
	assert.NoError(t, err)
}
