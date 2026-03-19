package auth_test

import (
	"testing"

	"github.com/uwr-tournament/server-go/internal/auth"
)

func TestPasswordHasher_Hash_GeneratesDifferentHashesForSamePassword(t *testing.T) {
	hasher := auth.NewPasswordHasher()
	password := "testPassword123"

	hash1, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Hash() failed: %v", err)
	}

	hash2, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Hash() failed: %v", err)
	}

	if hash1 == hash2 {
		t.Error("Hash() should generate different hashes for the same password")
	}
}

func TestPasswordHasher_Hash_EmptyPasswordError(t *testing.T) {
	hasher := auth.NewPasswordHasher()
	_, err := hasher.Hash("")
	if err == nil {
		t.Error("Hash() should return error for empty password")
	}
}

func TestPasswordHasher_Compare_MatchesCorrectPassword(t *testing.T) {
	hasher := auth.NewPasswordHasher()
	password := "testPassword123"

	hash, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Hash() failed: %v", err)
	}

	err = hasher.Compare(hash, password)
	if err != nil {
		t.Errorf("Compare() should match correct password: %v", err)
	}
}

func TestPasswordHasher_Compare_RejectsWrongPassword(t *testing.T) {
	hasher := auth.NewPasswordHasher()
	password := "testPassword123"
	wrongPassword := "wrongPassword456"

	hash, err := hasher.Hash(password)
	if err != nil {
		t.Fatalf("Hash() failed: %v", err)
	}

	err = hasher.Compare(hash, wrongPassword)
	if err == nil {
		t.Error("Compare() should reject wrong password")
	}
}

func TestPasswordHasher_GenerateToken_ReturnsRandomToken(t *testing.T) {
	hasher := auth.NewPasswordHasher()

	token1 := hasher.GenerateToken(20)
	token2 := hasher.GenerateToken(20)

	if len(token1) != 20 {
		t.Errorf("GenerateToken() should return token of length 20, got %d", len(token1))
	}

	if token1 == token2 {
		t.Error("GenerateToken() should return different tokens on each call")
	}
}
