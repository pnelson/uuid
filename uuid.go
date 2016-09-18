// Package uuid generates cryptographically strong pseudo-random UUIDv4's.
package uuid

import (
	"crypto/rand"
	"fmt"
)

// ByteSize is the number of bytes in a UUIDv4.
const ByteSize = 16

// Format is the UUIDv4 format specifier.
const Format = "%08x-%04x-%04x-%04x-%08x"

// New returns a new UUIDv4.
func New() string {
	b := make([]byte, ByteSize)
	_, err := rand.Read(b)
	if err != nil {
		// Cryptographic pseudo-random number generation shouldn't fail, but
		// if it does it is probably worth the panic.
		panic(fmt.Sprintf("uuid: %v", err))
	}
	return fmt.Sprintf(Format, b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}
