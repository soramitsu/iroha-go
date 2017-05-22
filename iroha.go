package iroha

import (
	"io"

	"github.com/upamune/ed25519"
)

type PublicKey []byte
type PrivateKey []byte

func NewKeyPair(rand io.Reader) (PublicKey, PrivateKey, error) {
	pub, pri, err := ed25519.GenerateKey(rand)
	return PublicKey(pub), PrivateKey(pri), err
}

func (priv PrivateKey) Sign(message []byte) []byte {
	return ed25519.Sign(ed25519.PrivateKey(priv), message)
}

func (pub PublicKey) Verify(message, sig []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(pub), message, sig)
}
