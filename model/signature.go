package model

import "github.com/google/flatbuffers/go"

type Signature struct {
	Pubkey []byte
	Sig []byte
}

func (s Signature) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return 0
}
