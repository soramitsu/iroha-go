package mock

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/model"
	"github.com/soramitsu/iroha-go/protocol"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type mockClientServiceServer struct{}

func (ss mockClientServiceServer) Torii(ctx context.Context, transaction *protocol.Transaction) (*flatbuffers.Builder, error) {
	res := model.Response{
		Message: "test_response_message",
		Signature: model.Signature{
			Pubkey: []byte("test_signature"),
			Sig:    []byte("test_signature"),
		},
		Code: model.OK,
	}
	builder := flatbuffers.NewBuilder(0)
	ro := res.Serialize(builder)
	builder.Finish(ro)

	return builder, nil
}

func NewMockServer() *grpc.Server {
	s := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))
	ss := mockClientServiceServer{}
	protocol.RegisterClientServiceServer(s, ss)
	return s
}
