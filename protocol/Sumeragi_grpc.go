//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: endpoint

package protocol

import "github.com/google/flatbuffers/go"

import (
  context "golang.org/x/net/context"
  grpc "google.golang.org/grpc"
)

// Client API for Sumeragi service
type SumeragiClient interface{
  Communicate(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* SumeragiResponse, error)  
}

type sumeragiClient struct {
  cc *grpc.ClientConn
}

func NewSumeragiClient(cc *grpc.ClientConn) SumeragiClient {
  return &sumeragiClient{cc}
}

func (c *sumeragiClient) Communicate(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* SumeragiResponse, error) {
  out := new(SumeragiResponse)
  err := grpc.Invoke(ctx, "/protocol.Sumeragi/Communicate", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for Sumeragi service
type SumeragiServer interface {
  Communicate(context.Context, *Block) (*flatbuffers.Builder, error)  
}

func RegisterSumeragiServer(s *grpc.Server, srv SumeragiServer) {
  s.RegisterService(&_Sumeragi_serviceDesc, srv)
}

func _Sumeragi_Communicate_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(Block)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(SumeragiServer).Communicate(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/protocol.Sumeragi/Communicate",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(SumeragiServer).Communicate(ctx, req.(* Block))
  }
  return interceptor(ctx, in, info, handler)
}


var _Sumeragi_serviceDesc = grpc.ServiceDesc{
  ServiceName: "protocol.Sumeragi",
  HandlerType: (*SumeragiServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Communicate",
      Handler: _Sumeragi_Communicate_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
  },
}
