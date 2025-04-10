package helpers

import (
	"crypto/sha256"
	"crypto/subtle"
)

func Sha256Sum(s string) []byte {
	sum := sha256.Sum256([]byte(s))
	arr := make([]byte, len(sum))
	copy(arr, sum[:])

	return arr
}

// secureCompare calculates sha256 hash of parameters a and b and does constant time comparison
// to avoid time based attacks.
func SecureCompare(a, b string) int {
	aSum := Sha256Sum(a)
	bSum := Sha256Sum(b)

	return subtle.ConstantTimeCompare(aSum, bSum)
}
