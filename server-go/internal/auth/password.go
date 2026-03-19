package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher handles password hashing and verification using bcrypt
type PasswordHasher struct {
	cost int
}

// NewPasswordHasher creates a new PasswordHasher with default bcrypt cost
func NewPasswordHasher() PasswordHasherInterface {
	return &PasswordHasher{
		cost: bcrypt.DefaultCost,
	}
}

// Hash generates a bcrypt hash of the password
func (ph *PasswordHasher) Hash(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), ph.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Compare checks if the password matches the hash
func (ph *PasswordHasher) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// GenerateToken generates a random token of specified length
// This mimics Devise.friendly_token behavior from Rails
func (ph *PasswordHasher) GenerateToken(length int) string {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		// Fallback: use base64 random token
		return generateFallbackToken(length)
	}

	return base64.URLEncoding.EncodeToString(buffer)[:length]
}

// generateFallbackToken provides a fallback token generation
func generateFallbackToken(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	buffer := make([]byte, length)
	_, _ = rand.Read(buffer)
	for i := range buffer {
		buffer[i] = charset[buffer[i]%byte(len(charset))]
	}
	return string(buffer)
}
