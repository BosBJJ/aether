package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)


func GenerateKey(size int) ([]byte, error) {
	key := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, key)
	return key, err
}

func GenerateAES256Key() ([]byte, error) {
	return GenerateKey(32)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}