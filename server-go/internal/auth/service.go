package auth

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/uwr-tournament/server-go/internal/models"
	"github.com/uwr-tournament/server-go/internal/repositories"
)

// authServiceImpl handles authentication operations including OAuth and username+password
type authServiceImpl struct {
	userRepo       repositories.UserRepository
	passwordHasher PasswordHasherInterface
}

// NewAuthService creates a new AuthService with dependency injection
func NewAuthService(userRepo repositories.UserRepository, passwordHasher PasswordHasherInterface) AuthService {
	return &authServiceImpl{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

// AuthenticateWithPassword authenticates a user with email and password
func (as *authServiceImpl) AuthenticateWithPassword(email, password string) (*User, error) {
	// Validate inputs
	if err := as.validateEmail(email); err != nil {
		return nil, fmt.Errorf("invalid email: %w", err)
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	// Find user by email
	modelUser, err := as.userRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if modelUser == nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	if err := as.passwordHasher.Compare(modelUser.EncryptedPassword, password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Update sign-in tracking
	now := time.Now()
	modelUser.SignInCount++
	modelUser.LastSignInAt = modelUser.CurrentSignInAt
	modelUser.CurrentSignInAt = &now

	// Save updated user
	if err := as.userRepo.Update(modelUser); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return modelUserToAuthUser(modelUser), nil
}

// GetOrCreateFromOAuth implements the OAuth flow, finding or creating a user by email
func (as *authServiceImpl) GetOrCreateFromOAuth(oauthData *OAuthData) (*User, error) {
	if oauthData == nil {
		return nil, errors.New("oauth data cannot be nil")
	}

	if err := as.validateEmail(oauthData.Email); err != nil {
		return nil, fmt.Errorf("invalid provider email: %w", err)
	}

	// Find existing user by email
	modelUser, err := as.userRepo.GetByEmail(oauthData.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// Return existing user if found
	if modelUser != nil {
		return modelUserToAuthUser(modelUser), nil
	}

	// Create new user
	newUser := &models.User{
		Email:             oauthData.Email,
		Name:              oauthData.Name,
		Provider:          oauthData.Provider,
		EncryptedPassword: as.passwordHasher.GenerateToken(20),
		SignInCount:       0,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := as.userRepo.Create(newUser); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return modelUserToAuthUser(newUser), nil
}

// validateEmail checks if an email is valid
func (as *authServiceImpl) validateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	// Simple email regex validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

// validatePassword checks if a password meets minimum requirements
func (as *authServiceImpl) validatePassword(password string) error {
	if password == "" {
		return errors.New("password cannot be empty")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

// Register creates a new user with username and password
func (as *authServiceImpl) Register(email, name, password string) (*User, error) {
	// Validate inputs
	if err := as.validateEmail(email); err != nil {
		return nil, fmt.Errorf("invalid email: %w", err)
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if err := as.validatePassword(password); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	// Check if user already exists
	existing, err := as.userRepo.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existing != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := as.passwordHasher.Hash(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	newUser := &models.User{
		Email:             email,
		Name:              name,
		EncryptedPassword: hashedPassword,
		SignInCount:       0,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	if err := as.userRepo.Create(newUser); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return modelUserToAuthUser(newUser), nil
}

// modelUserToAuthUser converts a models.User to an auth.User
func modelUserToAuthUser(modelUser *models.User) *User {
	if modelUser == nil {
		return nil
	}
	return &User{
		ID:                  modelUser.ID,
		Email:               modelUser.Email,
		Name:                modelUser.Name,
		EncryptedPassword:   modelUser.EncryptedPassword,
		Provider:            modelUser.Provider,
		ResetPasswordToken:  modelUser.ResetPasswordToken,
		ResetPasswordSentAt: modelUser.ResetPasswordSentAt,
		RememberCreatedAt:   modelUser.RememberCreatedAt,
		SignInCount:         modelUser.SignInCount,
		CurrentSignInAt:     modelUser.CurrentSignInAt,
		LastSignInAt:        modelUser.LastSignInAt,
		CurrentSignInIP:     modelUser.CurrentSignInIP,
		LastSignInIP:        modelUser.LastSignInIP,
		CreatedAt:           modelUser.CreatedAt,
		UpdatedAt:           modelUser.UpdatedAt,
	}
}
