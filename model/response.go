package model

import (
	"github.com/soramitsu/iroha-go/protocol"
	"github.com/google/flatbuffers/go"
)

type Response struct {
	Code      ResponseCode
	Message   string
	Signature Signature
}

func (r Response) Serialize(builder *flatbuffers.Builder)  flatbuffers.UOffsetT {
	msg := builder.CreateString(r.Message)
	sig := r.Signature.Serialize(builder)

	protocol.SumeragiResponseStart(builder)
	protocol.SumeragiResponseAddMessage(builder, msg)
	protocol.SumeragiResponseAddSignature(builder, sig)
	protocol.SumeragiResponseAddCode(builder, r.Code.ToByte())

	return protocol.SumeragiResponseEnd(builder)
}

func NewResponse(res *protocol.SumeragiResponse) *Response {
	r := &Response{
		Code: ResponseCode(res.Code()),
		Message: string(res.Message()),
	}

	sig := protocol.Signature{}
	res.Signature(&sig)

	r.Signature = Signature{
		Pubkey: sig.PubkeyBytes(),
		Sig: sig.SigBytes(),
	}

	return r
}


//go:generate stringer -type=ResponseCode
type ResponseCode int

func (rc ResponseCode) ToByte() byte {
	return byte(rc)
}

const (
	OK ResponseCode = iota + 1
	Fail
)
