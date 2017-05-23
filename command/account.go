package command

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
	"github.com/soramitsu/iroha-go/model"
)

type AddAccount struct {
	model.Account
}

func (cmd AddAccount) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	account := cmd.Account.Serialize(builder)
	iroha.AccountAddStart(builder)
	iroha.AccountAddAddAccount(builder, account)
	return iroha.AccountAddEnd(builder)
}

func (cmd AddAccount) Type() byte {
	return iroha.CommandAccountAdd
}
