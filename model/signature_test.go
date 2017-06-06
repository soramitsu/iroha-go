package model

import (
	"testing"

	"github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
)

func TestSignature(t *testing.T) {
	sig := Signature{
		Pubkey: []byte("this_is_test_pubkey"),
		Sig:    []byte("this_is_test_sig"),
	}

	builder := flatbuffers.NewBuilder(0)
	sigOffset := sig.Serialize(builder)
	builder.Finish(sigOffset)

	b := builder.FinishedBytes()

	sig2 := Signature{}
	sig2.Deserialize(b, 0)

	assert.Equal(t, sig, sig2)
}
