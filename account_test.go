package iroha

import (
	"testing"

	"github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountSerialize(t *testing.T) {
	account := Account{
		UserName:    "test_username",
		PubKey:      "test_public_key",
		Quorum:      uint16(5),
		Signatories: []string{"signatory_1", "signatory_2", "signatory_3"},
	}

	builder := flatbuffers.NewBuilder(0)
	accountOffset := account.Serialize(builder)
	builder.Finish(accountOffset)
	bytes := builder.FinishedBytes()

	account2 := Account{}
	require.NoError(t, account2.Deserialize(bytes, 0))

	assert.Equal(t, account, account2)
}
