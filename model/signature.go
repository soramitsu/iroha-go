package iroha

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
)

type Signature struct {
	PublicKey string
	Signature string
	Timestamp uint64
}

func (s Signature) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	pubkey := builder.CreateString(s.PublicKey)
	sig := builder.CreateString(s.Signature)

	iroha.SignatureStart(builder)
	iroha.SignatureAddPublicKey(builder, pubkey)
	iroha.SignatureAddSignature(builder, sig)
	iroha.SignatureAddTimestamp(builder, s.Timestamp)
	return iroha.SignatureEnd(builder)
}
