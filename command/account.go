package command

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
	"github.com/soramitsu/iroha-go/model"
)

type AddAccount struct {
	model.Account
}

func (cmd *AddAccount) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	cmdBuilder := flatbuffers.NewBuilder(0)
	account := cmd.Account.Serialize(cmdBuilder)
	cmdBuilder.Finish(account)
	b := cmdBuilder.FinishedBytes()

	bv := builder.CreateByteVector(b)
	iroha.AccountAddStart(builder)
	iroha.AccountAddAddAccount(builder, bv)
	return iroha.AccountAddEnd(builder)
}

func (cmd *AddAccount) Deserialize(table *flatbuffers.Table) {
	var icmd iroha.AccountAdd
	icmd.Init(table.Bytes, table.Pos)
	b := icmd.AccountBytes()

	account := model.Account{}
	account.Deserialize(b, 0)
	cmd.Account = account
}

func (cmd *AddAccount) Type() byte {
	return iroha.CommandAccountAdd
}

type RemoveAccount struct {
	model.Account
}

func (cmd *RemoveAccount) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	pubkey := builder.CreateString(cmd.Account.PubKey)
	iroha.AccountRemoveStart(builder)
	iroha.AccountRemoveAddPubkey(builder, pubkey)
	return iroha.AccountRemoveEnd(builder)
}

func (cmd *RemoveAccount) Deserialize(table *flatbuffers.Table) {
	var icmd iroha.AccountRemove
	icmd.Init(table.Bytes, table.Pos)

	cmd.Account.PubKey = string(icmd.Pubkey())
}

func (cmd *RemoveAccount) Type() byte {
	return iroha.CommandAccountRemove
}

type AddSignatory struct {
	model.Account
}

func (cmd *AddSignatory) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	account := builder.CreateString(cmd.Account.PubKey)

	lsigs := len(cmd.Account.Signatories)
	ss := make([]flatbuffers.UOffsetT, lsigs)
	for i := 0; i < lsigs; i++ {
		ss[i] = builder.CreateString(cmd.Account.Signatories[i])
	}

	iroha.AccountAddSignatoryStartSignatoryVector(builder, lsigs)
	for i := lsigs - 1; i >= 0; i-- {
		builder.PrependUOffsetT(ss[i])
	}
	sigs := builder.EndVector(lsigs)

	iroha.AccountAddSignatoryStart(builder)
	iroha.AccountAddSignatoryAddAccount(builder, account)
	iroha.AccountAddSignatoryAddSignatory(builder, sigs)
	return iroha.AccountAddSignatoryEnd(builder)
}

func (cmd *AddSignatory) Deserialize(table *flatbuffers.Table) {
	var icmd iroha.AccountAddSignatory
	icmd.Init(table.Bytes, table.Pos)

	sigs := make([]string, icmd.SignatoryLength())
	for i := 0; i < icmd.SignatoryLength(); i++ {
		sigs[i] = string(icmd.Signatory(i))
	}
	cmd.Account.PubKey = string(icmd.Account())
	cmd.Account.Signatories = sigs
}

func (cmd *AddSignatory) Type() byte {
	return iroha.CommandAccountAddSignatory
}
