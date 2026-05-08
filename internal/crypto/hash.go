package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
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

func getHashInst(algorithm HashAlgorithm) (hash.Hash, error) {
	switch algorithm {
	case HashSHA1:
		return sha1.New(), nil
	case HashMD5:
		return md5.New(), nil
	case HashSHA3_256:
		return sha3.New256(), nil
	case HashSHA3_512:
		return sha3.New512(), nil
	case HashSHA256:
		return sha256.New(), nil
	case HashSHA512:
		return sha512.New(), nil
	default:
		return nil, fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}
}

func Hash(input string, algorithm HashAlgorithm) (string, error) {
	h, err := getHashInst(algorithm)
	if err != nil {
		return "", err
	}
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil)), nil 
}

func HashFile(inputPath string, algorithm HashAlgorithm) (string, error) {
	h, err := getHashInst(algorithm)
	if err != nil {
		return "", err
	}
	f, err := os.Open(inputPath)
	if err != nil {
		return "", fmt.Errorf("could not open file %q: %v", inputPath, err)
	}
	defer f.Close()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}