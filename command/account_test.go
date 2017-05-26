package command

import (
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
	"github.com/soramitsu/iroha-go/model"
	"github.com/stretchr/testify/assert"
)

func serialize(cmd Commander) []byte {
	builder := flatbuffers.NewBuilder(0)
	offset := cmd.Serialize(builder)
	builder.Finish(offset)
	return builder.FinishedBytes()
}

func TestAddAccount(t *testing.T) {
	cmd := &AddAccount{
		Account: model.Account{
			PubKey:      "account_public_key",
			Quorum:      uint16(5),
			Signatories: []string{"test_sig1", "test_sig2", "test_sig3"},
			UUID:        "",
			UserName:    "serizawa",
		},
	}

	buf := serialize(cmd)

	cmd2 := &AddAccount{}
	root := iroha.GetRootAsAccountAdd(buf, 0)
	table := root.Table()
	cmd2.Deserialize(&table)

	assert.Equal(t, cmd, cmd2)
}

func TestRemoveAccount(t *testing.T) {
	cmd := &RemoveAccount{
		PubKey: "account_public_key",
	}

	buf := serialize(cmd)

	cmd2 := &RemoveAccount{}
	root := iroha.GetRootAsAccountRemove(buf, 0)
	table := root.Table()
	cmd2.Deserialize(&table)

	assert.Equal(t, cmd, cmd2)
}

func TestAccountAddSignatory(t *testing.T) {
	cmd := &AccountAddSignatory{
		Account:     "account_public_key",
		Signatories: []string{"sig1", "sig2", "sig3"},
	}

	buf := serialize(cmd)

	cmd2 := &AccountAddSignatory{}
	root := iroha.GetRootAsAccountAddSignatory(buf, 0)
	table := root.Table()
	cmd2.Deserialize(&table)

	assert.Equal(t, cmd, cmd2)
}
