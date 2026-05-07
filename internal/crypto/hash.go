package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"fmt"
	"hash"
)

type HashAlgorithm string

const (
	HashSHA1 		HashAlgorithm = "sha1"
	HashMD5 		HashAlgorithm = "md5"
	HashSHA3_256 	HashAlgorithm = "sha3-256"
	HashSHA3_512 	HashAlgorithm = "sha3-512"
	HashSHA256 		HashAlgorithm = "sha256"
	HashSHA512 		HashAlgorithm = "sha512"
)

func Hash(input string, algorithm HashAlgorithm) (string, error) {
	var h hash.Hash
	switch algorithm {
	case HashSHA1:
		h = sha1.New()
	case HashMD5:
		h = md5.New()
	case HashSHA3_256:
		h = sha3.New256()
	case HashSHA3_512:
		h = sha3.New512()
	case HashSHA256:
		h = sha256.New()
	case HashSHA512:
		h = sha512.New()
	default:
		return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}
	if h == nil {
		return "", fmt.Errorf("failed to create hash for algorithm: %s", algorithm)
	}

	h.Write([]byte(input))
	return fmt.Sprintf("%x\n", h.Sum(nil)), nil 
}