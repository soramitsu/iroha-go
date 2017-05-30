package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
)

type Response struct {
	Code      byte
	Message   string
	Signature Signature
}

func (r Response) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	msg := builder.CreateString(r.Message)
	sig := r.Signature.Serialize(builder)

	iroha.ResponseStart(builder)
	iroha.ResponseAddMessage(builder, msg)
	iroha.ResponseAddSignature(builder, sig)
	iroha.ResponseAddCode(builder, r.Code)
	return iroha.ResponseEnd(builder)
}

func NewResponse(resp *iroha.Response) *Response {
	r := &Response{
		Code:resp.Code(),
		Message:string(resp.Message()),

	}

	sig := iroha.Signature{}
	resp.Signature(&sig)
	r.Signature = Signature{
		PublicKey: string(sig.PublicKey()),
		Signature: string(sig.SignatureBytes()),
		Timestamp: sig.Timestamp(),
	}

	return r
}
