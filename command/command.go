package command

import "github.com/google/flatbuffers/go"

type Commander interface {
	Serialize(*flatbuffers.Builder) flatbuffers.UOffsetT
	Deserialize(*flatbuffers.Table)
	Type() byte
}
