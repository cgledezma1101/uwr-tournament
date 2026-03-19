package auth

//go:generate mockgen -source=types.go -destination=mocks/mock_types.go -package=mocks

import "time"

// User represents an authenticated user in the system
type User struct {
	ID                  int
	Email               string
	Name                string
	EncryptedPassword   string
	Provider            string
	ResetPasswordToken  string
	ResetPasswordSentAt *time.Time
	RememberCreatedAt   *time.Time
	SignInCount         int
	CurrentSignInAt     *time.Time
	LastSignInAt        *time.Time
	CurrentSignInIP     string
	LastSignInIP        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

// OAuthData represents data from an OAuth provider
type OAuthData struct {
	Email    string
	Name     string
	Provider string
	UID      string
}

// PasswordHasherInterface defines password hashing operations
type PasswordHasherInterface interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
	GenerateToken(length int) string
}

type AuthService interface  {
	AuthenticateWithPassword(email, password string) (*User, error)
	Register(email, name, password string) (*User, error)
	GetOrCreateFromOAuth(oauthData *OAuthData) (*User, error)
}

// OAuthProvider defines the interface for OAuth providers
type OAuthProvider interface {
	GetAuthorizationURL(state string) string
	ExchangeCodeForToken(code string) (string, error)
	GetUserInfo(accessToken string) (*OAuthData, error)
}
