package crypto

import (
	"crypto/ed25519"
	"errors"
	"github.com/itsabgr/go-handy"
)

var ErrInvalidSig = errors.New("invalid sig")

func Verify(pk, msg, sig []byte) error {
	if !ed25519.Verify(pk, msg, sig) {
		return ErrInvalidSig
	}
	return nil
}

func Sign(sk, msg []byte) []byte {
	return ed25519.Sign(sk, msg)
}

func GenSK() []byte {
	_, sk, err := ed25519.GenerateKey(nil)
	handy.Throw(err)
	return sk
}

func DerivePK(b []byte) []byte {
	return b[32:]
}
