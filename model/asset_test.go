package model

import (
	"testing"

	"github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
)

func TestCurrency(t *testing.T) {
	c1 := Currency{
		Amount:    []byte("1234567"),
		Precision: uint8(3),
	}

	builder := flatbuffers.NewBuilder(0)
	c := c1.Serialize(builder)
	builder.Finish(c)
	b := builder.FinishedBytes()

	c2 := Currency{}
	c2.Deserialize(b, 0)

	assert.Equal(t, c1, c2)
}
