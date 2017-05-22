package command

import "github.com/google/flatbuffers/go"

type Commander interface {
	Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT
	Type() byte
}
