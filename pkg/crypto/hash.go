package crypto

import "crypto/sha256"

func Hash(b []byte) []byte {
	h := sha256.Sum256(b)
	return h[:]
}
