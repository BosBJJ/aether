package crypto

import (
    "testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
    password := "hunter2"

    hash, err := HashPassword(password)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if hash == "" {
        t.Fatal("expected a non-empty hash")
    }

    match, err := VerifyPassword(password, hash)
    if err != nil {
        t.Fatalf("unexpected error during comparison: %v", err)
    }
    if !match {
        t.Error("expected password to match hash")
    }
}

func TestGeneratePassword(t *testing.T) {
	password, err := GeneratePassword(15, false, false)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(password) != 15 {
		t.Errorf("expected length 15, got %d", len(password))
	}

	password2, _ := GeneratePassword(16, false, false)
    if password == password2 {
        t.Error("two generated passwords should not be identical")
    }
}



