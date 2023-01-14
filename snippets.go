package snippets

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
)

// Replace the first element of a slice with the last one and return a slice with the tail cut off
func removeSliceZero(s []int) []int {
	s[0] = s[len(s)-1]
	return s[:len(s)-1]
}

// Simultaneous calls for new random slices may overlap; time based seed is insufficient;
// Cryptographically secure approach to ensure unique slices
func crypto_seed() int64 {
	// https://stackoverflow.com/a/54491783

	var b [8]byte

	if _, err := crypto_rand.Read(b[:]); err != nil {
		panic("Cannot seed")
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}
