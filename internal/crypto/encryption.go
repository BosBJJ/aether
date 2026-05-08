package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)
func generateNonce(size int) ([]byte, error) {
	nonce := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}


func EncryptAES(plaintext, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("unable to create new AES cipher: %v", err)
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("unable to create GCM: %v", err)
	}
	nonce, err := generateNonce(aesGCM.NonceSize())
	if err != nil {
		return "", fmt.Errorf("unable to create nonce: %v", err)
	}
	cipherText := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return Base64Encode(cipherText), nil
}

func DecryptAES(s string, key []byte) ([]byte, error){
	cipherText, err := Base64Decode(s)
	if err != nil {
		return nil, fmt.Errorf("could not convert string to []byte: %v", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("unable to create new AES cipher: %v", err)
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("unable to create GCM: %v", err)
	}
	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, actualCipher := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, actualCipher, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %v", err)
	}
	return plaintext, nil
}

func EncryptFileAES(inputPath, outputPath string, key []byte) error {
	plaintext, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("could not read file %q: %v", inputPath, err)
	}
	encrypted, err := EncryptAES(plaintext, key)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputPath, []byte(encrypted), 0600)
	if err != nil {
		return fmt.Errorf("unable to write to file: %v" ,err)
	}
	return nil
}

func DecryptFileAES(inputPath, outputPath string, key []byte) error {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("could not read file %q: %v", inputPath, err)
	}
	plaintext, err := DecryptAES(string(data), key)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputPath, plaintext, 0600)
	if err != nil {
		return fmt.Errorf("could not write to file: %v", err)
	}
	return nil
}


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