package mock

import (
	"time"

	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/protocol"
	"github.com/soramitsu/iroha-go/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type mockSumeragiServer struct{}

func (ss mockSumeragiServer) Torii(ctx context.Context, ev *protocol.Transaction) (*flatbuffers.Builder, error) {
	res := model.Response{
		Message: "test_response_message",
		Signature: model.Signature{
			PublicKey: "test_signature",
			Signature: "test_signature",
			Timestamp: uint64(time.Now().Unix()),
		},
		Code: protocol.CodeUNDECIDED,
	}
	builder := flatbuffers.NewBuilder(0)
	ro := res.Serialize(builder)
	builder.Finish(ro)

	return builder, nil
}

func (ss mockSumeragiServer) Verify(ctx context.Context, ev *protocol.ConsensusEvent) (*flatbuffers.Builder, error) {
	builder := flatbuffers.NewBuilder(0)
	builder.Finish(0)

	return builder, nil
}

func NewMockServer() *grpc.Server {
	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	ss := mockSumeragiServer{}
	protocol.RegisterSumeragiServer(s, ss)
	return s
}
