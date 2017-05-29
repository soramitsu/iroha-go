package mock

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type mockSumeragiServer struct{}

func (ss mockSumeragiServer) Torii(ctx context.Context, ev *iroha.Transaction) (*flatbuffers.Builder, error) {
	builder := flatbuffers.NewBuilder(0)
	builder.Finish(0)
	// TODO (@upamune): return iroha.Transaction as *flatbuffers.Builder
	return builder, nil
}

func (ss mockSumeragiServer) Verify(ctx context.Context, ev *iroha.ConsensusEvent) (*flatbuffers.Builder, error) {
	builder := flatbuffers.NewBuilder(0)
	builder.Finish(0)

	return builder, nil
}

func NewMockServer() *grpc.Server {
	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	ss := mockSumeragiServer{}
	iroha.RegisterSumeragiServer(s, ss)
	return s
}
