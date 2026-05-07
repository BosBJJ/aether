package crypto

import "testing"

func TestHash(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		algorithm HashAlgorithm
		wantErr   bool
	}{
		{"SHA1", "test", HashSHA1, false},
		{"MD5", "hello", HashMD5, false},
		{"SHA256", "hello", HashSHA256, false},
		{"SHA512", "password123", HashSHA512, false},
		{"SHA3-256", "test", HashSHA3_256, false},
		{"SHA512", "hello world", HashSHA512, false},
		{"Invalid", "hello", "invalid-alg", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := Hash(tc.input, tc.algorithm)

			if tc.wantErr {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if hash == "" {
				t.Error("hash should not be empty")
			}

			// Optional: print for visual confirmation
			t.Logf("%-10s → %s", tc.algorithm, hash)
		})
	}
}