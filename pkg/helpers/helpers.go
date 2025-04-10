package helpers

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"
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

func b64Decode(s string) ([]byte, error) {
	s = strings.ReplaceAll(s, ".", "+")
	return base64.RawStdEncoding.DecodeString(s)
}

func HashStr(str string) string {
	salt, err := b64Decode("s1DxLLMBt3CuPs4F7JeNfQ0yrVGF5iPZ")
	iteration := 10000
	if err != nil {
		fmt.Println("Failed to base64 decode the salt:", err)
	}

	key, err := pbkdf2.Key(sha256.New, str, salt, iteration, 32)

	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", key)
}
