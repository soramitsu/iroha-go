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
	cmdBuilder := flatbuffers.NewBuilder(0)

	account := cmd.Account.Serialize(cmdBuilder)

	cmdBuilder.Finish(account)
	b := cmdBuilder.FinishedBytes()

	bv := builder.CreateByteVector(b)
	iroha.AccountAddStart(builder)
	iroha.AccountAddAddAccount(builder, bv)
	return iroha.AccountAddEnd(builder)
}

func (cmd AddAccount) Type() byte {
	return iroha.CommandAccountAdd
}
