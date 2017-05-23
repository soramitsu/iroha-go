package iroha

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/command"
	"github.com/soramitsu/iroha-go/iroha"
)

type Transaction struct {
	Command    command.Commander
	PublicKey  string
	Signatures []Signature
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

func (t *Transaction) Deserialize(buf []byte, offset flatbuffers.UOffsetT) {
	transaction := iroha.GetRootAsTransaction(buf, offset)

	// Signatures
	lsig := transaction.SignaturesLength()
	sigs := make([]Signature, lsig)
	for i := 0; i < lsig; i++ {
		s := iroha.Signature{}
		transaction.Signatures(&s, i)
		sigs[i] = Signature{
			PublicKey: string(s.PublicKey()),
			Signature: string(s.SignatureBytes()),
			Timestamp: s.Timestamp(),
		}
	}

	// TODO: Deserialize Command
	//var cmd flatbuffers.Table
	//pp.Println(transaction.Command(&cmd))
	//aa := iroha.GetRootAsAccountAdd(cmd.Bytes, 0).Table()
	//pp.Println(aa.Bytes)

	t.Signatures = sigs
	t.PublicKey = string(transaction.CreatorPubKey())
	t.Timestamp = transaction.Timestamp()
}
