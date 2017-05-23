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
	builder2 := flatbuffers.NewBuilder(0)

	account := cmd.Account.Serialize(builder2)

	builder2.Finish(account)
	b := builder2.FinishedBytes()

	bv := builder.CreateByteVector(b)
	iroha.AccountAddStart(builder)
	iroha.AccountAddAddAccount(builder, bv)
	return iroha.AccountAddEnd(builder)
}

func (cmd AddAccount) Type() byte {
	return iroha.CommandAccountAdd
}
