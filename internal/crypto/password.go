package crypto

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/alexedwards/argon2id"
	"github.com/ccojocar/zxcvbn-go"
)

var DefaultParams = &argon2id.Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 4,
	SaltLength:  16,
	KeyLength:   32,
}


func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, DefaultParams)
	if err != nil {
		return "", fmt.Errorf("Could not hash password: %w", err)
	}
	return hash, nil
}

func VerifyPassword(password, hash string) (bool, error) {
	isSame, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, fmt.Errorf("Unable to compare password and hash: %w", err)
	}
	return isSame, nil
}

const (
    lowerChars   = "abcdefghijklmnopqrstuvwxyz"
    upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    digitChars   = "0123456789"
    specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func GeneratePassword(length int, hasSpecial, hasUppercase bool) (string, error) {
	result := make([]byte, length)
	charSet := lowerChars + digitChars
	if hasUppercase {
		charSet = charSet + upperChars
	}
	if hasSpecial {
		charSet = charSet + specialChars
	}
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		result[i] = charSet[num.Int64()]
	}

	return string(result), nil
}

func AnalyzePassword(password string) PasswordAnalysis {
	res := zxcvbn.PasswordStrength(password, nil)
	var matches []MatchResult
	for _, m := range res.MatchSequence {
		matches = append(matches, MatchResult{
			Pattern: m.Pattern,
			Token: m.Token,
			Entropy: m.Entropy,
		})
	}
	return PasswordAnalysis{
		Entropy: res.Entropy,
		CrackTime: res.CrackTimeDisplay,
		MatchSequence: matches,
	}
}

type PasswordAnalysis struct {
	Entropy 		float64 			`json:"entropy"`
	CrackTime		string				`json:"crack_time"`
	MatchSequence 	[]MatchResult 		`json:"match_sequence"` 
}

type MatchResult struct {
	Pattern 		string 				`json:"pattern"`
	Token 			string 				`json:"token"`
	Entropy			float64				`json:"entropy"`
}