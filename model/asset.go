package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/soramitsu/iroha-go/protocol"
)

type AnyAsset interface {
	Serialize(flatbuffers.Builder) flatbuffers.UOffsetT
}

type Asset struct {
	Name  string
	UUID  string
	asset AnyAsset
}

// TODO (@upamune): implement ComplexAsset
type ComplexAsset struct {
	Prop []byte
}

type Currency struct {
	Amount    []byte
	Precision uint8
}

func (c Currency) Serialize(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	amount := builder.CreateByteString(c.Amount)

	protocol.CurrencyStart(builder)
	protocol.CurrencyAddAmount(builder, amount)
	protocol.CurrencyAddPrecision(builder, c.Precision)
	return protocol.CurrencyEnd(builder)
}

func (c *Currency) Deserialize(b []byte, offset flatbuffers.UOffsetT) {
	currency := NewCurrency(protocol.GetRootAsCurrency(b, offset))

	c.Amount = currency.Amount
	c.Precision = currency.Precision
}

func NewCurrency(c *protocol.Currency) *Currency {
	return &Currency{
		Amount:    c.AmountBytes(),
		Precision: c.Precision(),
	}
}
