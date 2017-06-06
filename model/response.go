package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/protocol"
)

type Response struct {
	Code      ResponseCode
	Message   string
	Signature Signature
}

func (r Response) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	msg := builder.CreateString(r.Message)
	sig := r.Signature.Serialize(builder)

	protocol.SumeragiResponseStart(builder)
	protocol.SumeragiResponseAddMessage(builder, msg)
	protocol.SumeragiResponseAddSignature(builder, sig)
	protocol.SumeragiResponseAddCode(builder, r.Code.ToByte())

	return protocol.SumeragiResponseEnd(builder)
}

func (r *Response) Deserialize(b []byte, offset flatbuffers.UOffsetT) {
	resp := protocol.GetRootAsSumeragiResponse(b, offset)

	res := NewResponse(resp)

	r.Code = res.Code
	r.Message = res.Message
	r.Signature = res.Signature
}

func NewResponse(res *protocol.SumeragiResponse) *Response {
	r := &Response{
		Code:    NewResponseCode(res.Code()),
		Message: string(res.Message()),
	}

	sig := protocol.Signature{}
	res.Signature(&sig)

	r.Signature = Signature{
		Pubkey: sig.PubkeyBytes(),
		Sig:    sig.SigBytes(),
	}

	return r
}

//go:generate stringer -type=ResponseCode
type ResponseCode int

func NewResponseCode(code byte) ResponseCode {
	switch code {
	case protocol.ResponseCodeOK:
		return OK
	case protocol.ResponseCodeFAIL:
		return Fail
	}

	return Unknown
}

func (rc ResponseCode) ToByte() byte {
	switch rc {
	case OK:
		return protocol.ResponseCodeOK
	case Fail:
		return protocol.ResponseCodeFAIL
	}

	return 0
}

const (
	Unknown ResponseCode = iota
	OK
	Fail
)
