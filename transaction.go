package iroha

import (
	"github.com/google/flatbuffers/go"
	"github.com/pkg/errors"
	"github.com/soramitsu/iroha-go/command"
	"github.com/soramitsu/iroha-go/iroha"
	"github.com/soramitsu/iroha-go/model"
)

type Transaction struct {
	Command    command.Commander
	PublicKey  string
	Signatures []model.Signature
	Timestamp  uint64
}

func (t Transaction) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	pubkey := builder.CreateString(t.PublicKey)
	command := t.Command.Serialize(builder)

	lsigs := len(t.Signatures)
	ss := make([]flatbuffers.UOffsetT, lsigs)
	for i := 0; i < lsigs; i++ {
		ss[i] = t.Signatures[i].Serialize(builder)
	}

	iroha.TransactionStartSignaturesVector(builder, lsigs)
	for i := lsigs - 1; i >= 0; i-- {
		builder.PrependUOffsetT(ss[i])
	}
	sigs := builder.EndVector(lsigs)

	iroha.TransactionStart(builder)
	iroha.TransactionAddCommand(builder, command)
	iroha.TransactionAddCommandType(builder, t.Command.Type())
	iroha.TransactionAddCreatorPubKey(builder, pubkey)
	iroha.TransactionAddSignatures(builder, sigs)
	iroha.TransactionAddTimestamp(builder, t.Timestamp)

	return iroha.TransactionEnd(builder)
}

func (t *Transaction) Deserialize(buf []byte, offset flatbuffers.UOffsetT) error {
	transaction := iroha.GetRootAsTransaction(buf, offset)

	// Signatures
	lsig := transaction.SignaturesLength()
	sigs := make([]model.Signature, lsig)
	for i := 0; i < lsig; i++ {
		s := iroha.Signature{}
		transaction.Signatures(&s, i)
		sigs[i] = model.Signature{
			PublicKey: string(s.PublicKey()),
			Signature: string(s.SignatureBytes()),
			Timestamp: s.Timestamp(),
		}
	}

	var table flatbuffers.Table
	if ok := transaction.Command(&table); !ok {
		return errors.New("transaction.Command failed")
	}

	var cmd command.Commander
	switch transaction.CommandType() {
	case iroha.CommandAccountAdd:
		cmd = &command.AddAccount{}

	case iroha.CommandAccountRemove:
		cmd = &command.RemoveAccount{}

	case iroha.CommandAccountAddSignatory:
		cmd = &command.AccountAddSignatory{}

	default:
		return errors.New("unknown command type")
	}

	cmd.Deserialize(&table)
	t.Command = cmd

	t.Signatures = sigs
	t.PublicKey = string(transaction.CreatorPubKey())
	t.Timestamp = transaction.Timestamp()

	return nil
}
