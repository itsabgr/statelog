package utils

import "encoding/binary"

func Uint64ToBE(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}
