package iroha

import "testing"

type zeroReader struct{}

func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

func TestSignVerify(t *testing.T) {
	var zero zeroReader
	public, private, _ := NewKeyPair(zero)

	message := []byte("test message")
	sig := private.Sign(message)
	if !public.Verify(message, sig) {
		t.Errorf("valid signature rejected")
	}

	wrongMessage := []byte("wrong message")
	if public.Verify(wrongMessage, sig) {
		t.Errorf("signature of different message accepted")
	}
}

func BenchmarkNewKeyPair(b *testing.B) {
	var zero zeroReader
	for i := 0; i < b.N; i++ {
		if _, _, err := NewKeyPair(zero); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSigning(b *testing.B) {
	var zero zeroReader
	_, priv, err := NewKeyPair(zero)
	if err != nil {
		b.Fatal(err)
	}
	message := []byte("Hello, world!")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		priv.Sign(message)
	}
}

func BenchmarkVerification(b *testing.B) {
	var zero zeroReader
	pub, priv, err := NewKeyPair(zero)
	if err != nil {
		b.Fatal(err)
	}
	message := []byte("Hello, world!")
	signature := priv.Sign(message)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pub.Verify(message, signature)
	}
}
