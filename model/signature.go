package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/protocol"
)

type Signature struct {
	Pubkey []byte
	Sig    []byte
}

func (s Signature) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	pubkey := builder.CreateByteString(s.Pubkey)
	sig := builder.CreateByteString(s.Sig)

	protocol.SignatureStart(builder)
	protocol.SignatureAddPubkey(builder, pubkey)
	protocol.SignatureAddSig(builder, sig)
	return protocol.SignatureEnd(builder)
}

func (s *Signature) Deserialize(b []byte, offset flatbuffers.UOffsetT) {
	signature := protocol.GetRootAsSignature(b, offset)
	s.Pubkey = signature.PubkeyBytes()
	s.Sig = signature.SigBytes()
}
