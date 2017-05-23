package iroha

import (
	"testing"
	"time"

	"fmt"

	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/command"
	"github.com/soramitsu/iroha-go/model"
)

func TestTransaction_Serialize(t *testing.T) {
	account := model.Account{
		PubKey:      "account_public_key",
		Quorum:      uint16(5),
		Signatories: []string{"test_sig1", "test_sig2", "test_sig3"},
		UUID:        "this-is-dummy-uuid",
		UserName:    "serizawa",
	}
	addAccountCmd := command.AddAccount{
		Account: account,
	}

	lsigs := 3
	sigs := make([]Signature, lsigs)
	for i := 0; i < lsigs; i++ {
		s := Signature{
			PublicKey: fmt.Sprintf("test_public_key_%d", i+1),
			Signature: fmt.Sprintf("test_signature_%d", i+1),
			Timestamp: uint64(time.Now().Unix() + int64(i)),
		}
		sigs[i] = s
	}

	transaction := Transaction{
		Command:    addAccountCmd,
		Signatures: sigs,
		PublicKey:  "test_public_key",
		Timestamp:  uint64(time.Now().Unix()),
	}

	builder := flatbuffers.NewBuilder(0)
	tbuf := transaction.Serialize(builder)
	builder.Finish(tbuf)

	buf := builder.FinishedBytes()

	transaction2 := Transaction{}
	transaction2.Deserialize(buf, 0)

	//pp.Println(transaction)
	//pp.Println(transaction2)
	//assert.Equal(t, transaction, transaction2)
}
