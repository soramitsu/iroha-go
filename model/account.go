package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/iroha"
)

type Account struct {
	PubKey      string
	Quorum      uint16
	Signatories []string
	UUID        string
	UserName    string
}

func (a Account) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	pubkey := builder.CreateString(a.PubKey)
	alias := builder.CreateString(a.UserName)

	l := len(a.Signatories)

	ss := make([]flatbuffers.UOffsetT, l)
	for i := 0; i < l; i++ {
		ss[i] = builder.CreateString(a.Signatories[i])
	}

	iroha.AccountStartSignatoriesVector(builder, l)
	for i := l - 1; i >= 0; i-- {
		builder.PrependUOffsetT(ss[i])
	}
	signatories := builder.EndVector(l)

	iroha.AccountStart(builder)
	iroha.AccountAddAlias(builder, alias)
	iroha.AccountAddUseKeys(builder, a.Quorum)
	iroha.AccountAddPubKey(builder, pubkey)
	iroha.AccountAddSignatories(builder, signatories)

	return iroha.AccountEnd(builder)
}

func (a *Account) Deserialize(buf []byte, offset flatbuffers.UOffsetT) error {
	account := iroha.GetRootAsAccount(buf, offset)

	signatories := make([]string, account.SignatoriesLength())
	for i := 0; i < account.SignatoriesLength(); i++ {
		signatory := string(account.Signatories(i))
		signatories[i] = signatory
	}

	a.PubKey = string(account.PubKey())
	a.Quorum = account.UseKeys()
	a.UserName = string(account.Alias())
	a.Signatories = signatories

	return nil
}

type AccountsQuery struct {
	Domain      string
	IsCommitted bool
	Limit       int
	Offset      int
	Order       string
}

type AccountQuery struct {
	Domain      string
	IsCommitted bool
	UUID        string
	UserName    string
}
