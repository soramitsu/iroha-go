package mock

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type mockSumeragiServer struct{}

func (ss *mockSumeragiServer) Torii(ctx context.Context, ev *iroha.Transaction) (*flatbuffers.Builder, error) {
	// TODO (@upamune): return iroha.Transaction as *flatbuffers.Builder
	return nil, nil
}

func (ss *mockSumeragiServer) Verify(ctx context.Context, ev *iroha.ConsensusEvent) (*flatbuffers.Builder, error) {
	return nil, nil
}

func NewMockServer() *grpc.Server {
	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	ss := mockSumeragiServer{}
	iroha.RegisterSumeragiServer(s, ss)
	return s
}
