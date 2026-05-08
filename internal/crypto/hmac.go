package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
)



func GenerateHMAC(message, key []byte, algorithm HashAlgorithm) (string, error) {
	var h hash.Hash
	switch algorithm{
	case HashSHA1:
		h = hmac.New(sha1.New, key)
	case HashMD5:
		h = hmac.New(md5.New, key)
	case HashSHA3_256:
		h = hmac.New(func() hash.Hash { return sha3.New256()}, key)
	case HashSHA3_512:
		h = hmac.New(func() hash.Hash { return sha3.New512()}, key)
	case HashSHA256:
		h = hmac.New(sha256.New, key)
	case HashSHA512:
		h = hmac.New(sha512.New, key)
	default:
		return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)	
	}
	h.Write(message)
	signature := h.Sum(nil)

	return hex.EncodeToString(signature), nil
} 

func VerifyHMAC(message, key []byte, expected string, algorithm HashAlgorithm) (bool, error) {
	mac, err := GenerateHMAC(message, key, algorithm)
	if err != nil {
		return false, err
	}
	return hmac.Equal([]byte(mac), []byte(expected)), nil
}