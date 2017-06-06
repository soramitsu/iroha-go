package model

import (
	"testing"

	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/protocol"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	resp := Response{
		Code:      OK,
		Message:   "this_is_test_message",
		Signature: Signature{
			Pubkey: []byte("this_is_test_pubkey"),
			Sig:    []byte("this_is_test_signature"),
		},
	}

	builder := flatbuffers.NewBuilder(0)
	respOffset := resp.Serialize(builder)
	builder.Finish(respOffset)
	b := builder.FinishedBytes()

	resp2 := Response{}
	resp2.Deserialize(b, 0)

	assert.Equal(t, resp, resp2)
}

func TestNewResponseCode(t *testing.T) {
	patterns := []struct {
		Expected ResponseCode
		Test     byte
	}{
		{
			Unknown,
			byte(3),
		},
		{
			OK,
			protocol.ResponseCodeOK,
		},
		{
			Fail,
			protocol.ResponseCodeFAIL,
		},
	}

	for _, p := range patterns {
		rc := NewResponseCode(p.Test)
		assert.Equal(t, p.Expected, rc)
	}
}
