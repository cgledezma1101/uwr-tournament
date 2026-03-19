package auth_test

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/auth"
	authmocks "github.com/uwr-tournament/server-go/internal/auth/mocks"
	"github.com/uwr-tournament/server-go/internal/models"
	repomocks "github.com/uwr-tournament/server-go/internal/repositories/mocks"
	"go.uber.org/mock/gomock"
)

func TestAuthService_AuthenticateWithPassword(t *testing.T) {
	t.Run("SuccessfulLogin", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		user := &models.User{
			ID:                1,
			Email:             "test@example.com",
			Name:              "Test User",
			EncryptedPassword: "model_password",
		}
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(user.Email).Return(user, nil)
		repo.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().Compare(user.EncryptedPassword, "correctPassword123").Return(nil).Times(1)

		service := auth.NewAuthService(repo, hasher)

		authenticated, err := service.AuthenticateWithPassword(user.Email, "correctPassword123")
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(authenticated.ID).To(Equal(user.ID))
		g.Expect(authenticated.Email).To(Equal(user.Email))
		g.Expect(authenticated.Name).To(Equal(user.Name))
		g.Expect(authenticated.EncryptedPassword).To(Equal(user.EncryptedPassword))
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(gomock.Any()).Return(&models.User{}, nil).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(errors.New("invalid password")).Times(1)
		service := auth.NewAuthService(repo, hasher)

		authenticated, err := service.AuthenticateWithPassword("test@example.com", "wrongPassword")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("invalid email or password"))
		g.Expect(authenticated).To(BeNil())
	})

	t.Run("InvalidEmail", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		repo := repomocks.NewMockUserRepository(ctrl)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		service := auth.NewAuthService(repo, hasher)

		authenticated, err := service.AuthenticateWithPassword("invalid-email", "wrongPassword")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("invalid email: invalid email format"))
		g.Expect(authenticated).To(BeNil())
	})

	t.Run("UserNotFound", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(gomock.Any()).Return(nil, nil).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		service := auth.NewAuthService(repo, hasher)

		authenticated, err := service.AuthenticateWithPassword("nonexistent@example.com", "password")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("invalid email or password"))
		g.Expect(authenticated).To(BeNil())
	})

	t.Run("GetUserError", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(gomock.Any()).Return(nil, errors.New("database error")).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		service := auth.NewAuthService(repo, hasher)

		authenticated, err := service.AuthenticateWithPassword("nonexistent@example.com", "password")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("failed to find user: database error"))
		g.Expect(authenticated).To(BeNil())
	})

	t.Run("FailedToUpdateUser", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(gomock.Any()).Return(&models.User{}, nil).Times(1)
		repo.EXPECT().Update(gomock.Any()).Return(errors.New("update error")).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().Compare(gomock.Any(), gomock.Any()).Return(nil).Times(1)
		service := auth.NewAuthService(repo, hasher)

		authenticated, err := service.AuthenticateWithPassword("nonexistent@example.com", "password")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("failed to save user: update error"))
		g.Expect(authenticated).To(BeNil())
	})
}

func TestAuthService_GetOrCreateFromOAuth(t *testing.T) {
	t.Run("NoAuthData", func(t *testing.T) {
		g := NewWithT(t)
		service := auth.NewAuthService(nil, nil)

		user, err := service.GetOrCreateFromOAuth(nil)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("oauth data cannot be nil"))
		g.Expect(user).To(BeNil())
	})

	t.Run("InvalidEmail", func(t *testing.T) {
		g := NewWithT(t)
		service := auth.NewAuthService(nil, nil)

		oauthData := &auth.OAuthData{
			Email:    "invalid-email",
			Name:     "Test User",
			Provider: "google",
		}

		user, err := service.GetOrCreateFromOAuth(oauthData)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("invalid provider email"))
		g.Expect(user).To(BeNil())
	})

	t.Run("ErrorGettingUser", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(gomock.Any()).Return(nil, errors.New("database error")).Times(1)

		service := auth.NewAuthService(repo, nil)

		oauthData := &auth.OAuthData{
			Email:    "test@example.com",
			Name:     "Test User",
			Provider: "google",
		}

		user, err := service.GetOrCreateFromOAuth(oauthData)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("failed to find user: database error"))
		g.Expect(user).To(BeNil())
	})

	t.Run("ExistingUser", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		repo := repomocks.NewMockUserRepository(ctrl)

		existingUser := &models.User{
			ID:    1,
			Email: "oauth@example.com",
			Name:  "Existing Name",
		}

		oauthData := &auth.OAuthData{
			Email:    "oauth@example.com",
			Name:     "OAuth Name",
			Provider: "google",
		}
		repo.EXPECT().GetByEmail(oauthData.Email).Return(existingUser, nil).Times(1)

		service := auth.NewAuthService(repo, authmocks.NewMockPasswordHasherInterface(ctrl))

		user, err := service.GetOrCreateFromOAuth(oauthData)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(user.ID).To(Equal(existingUser.ID))
		g.Expect(user.Email).To(Equal(existingUser.Email))
		g.Expect(user.Name).To(Equal(existingUser.Name))
	})

	t.Run("NewUser", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().GenerateToken(20).Return("random-hash").Times(1)

		repo := repomocks.NewMockUserRepository(ctrl)

		oauthData := &auth.OAuthData{
			Email:    "newoauth@example.com",
			Name:     "New OAuth User",
			Provider: "facebook",
		}
		repo.EXPECT().GetByEmail(oauthData.Email).Return(nil, nil).Times(1)
		repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(user *models.User) error {
			g.Expect(user.Email).To(Equal(oauthData.Email))
			g.Expect(user.Name).To(Equal(oauthData.Name))
			g.Expect(user.Provider).To(Equal(oauthData.Provider))
			g.Expect(user.EncryptedPassword).To(Equal("random-hash"))
			return nil
		}).Times(1)

		service := auth.NewAuthService(repo, hasher)

		user, err := service.GetOrCreateFromOAuth(oauthData)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(user.Email).To(Equal(oauthData.Email))
		g.Expect(user.Name).To(Equal(oauthData.Name))
		g.Expect(user.Provider).To(Equal(oauthData.Provider))
		g.Expect(user.EncryptedPassword).To(Equal("random-hash"))
	})

	t.Run("NewUser_FailedCreation", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().GenerateToken(20).Return("random-hash").Times(1)

		repo := repomocks.NewMockUserRepository(ctrl)

		oauthData := &auth.OAuthData{
			Email:    "newoauth@example.com",
			Name:     "New OAuth User",
			Provider: "facebook",
		}
		repo.EXPECT().GetByEmail(oauthData.Email).Return(nil, nil).Times(1)
		repo.EXPECT().Create(gomock.Any()).Return(errors.New("create error")).Times(1)

		service := auth.NewAuthService(repo, hasher)

		user, err := service.GetOrCreateFromOAuth(oauthData)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("failed to create user: create error"))
		g.Expect(user).To(BeNil())
	})
}

func TestAuthService_Register(t *testing.T) {
	t.Run("SuccessfulRegistration", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		email := "newuser@example.com"
		hashedPassword := "hashed-password"
		username := "New User"
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(email).Return(nil, nil).Times(1)
		repo.EXPECT().Create(gomock.Any()).DoAndReturn(func(user *models.User) error {
			g.Expect(user.Email).To(Equal(email))
			g.Expect(user.Name).To(Equal(username))
			g.Expect(user.Provider).To(Equal(""))
			g.Expect(user.EncryptedPassword).To(Equal(hashedPassword))
			return nil
		}).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().Hash("password123").Return(hashedPassword, nil).Times(1)
		service := auth.NewAuthService(repo, hasher)

		user, err := service.Register(email, username, "password123")
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(user.Email).To(Equal(email))
		g.Expect(user.Name).To(Equal(username))
		g.Expect(user.EncryptedPassword).To(Equal(hashedPassword))
		g.Expect(user.Provider).To(BeEmpty())
	})

	t.Run("InvalidEmail", func(t *testing.T) {
		g := NewWithT(t)

		service := auth.NewAuthService(nil, nil)

		user, err := service.Register("invalid-email", "User", "password123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("invalid email: invalid email format"))
		g.Expect(user).To(BeNil())
	})

	t.Run("EmptyName", func(t *testing.T) {
		g := NewWithT(t)

		service := auth.NewAuthService(nil, nil)

		user, err := service.Register("valid@email.com", "", "password123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("name cannot be empty"))
		g.Expect(user).To(BeNil())
	})

	invalidPasswords := []struct {
		Name     string
		Password string
		Error string
	}{
		{"TooShort", "short", "password must be at least 8 characters"},
		{"Empty", "", "password cannot be empty"},
	}
	for _, tc := range invalidPasswords {
		t.Run("InvalidPassword_"+tc.Name, func(t *testing.T) {
			g := NewWithT(t)

			service := auth.NewAuthService(nil, nil)

			user, err := service.Register("valid@email.com", "Some user", tc.Password)
			g.Expect(err).To(HaveOccurred())
			g.Expect(err.Error()).To(ContainSubstring("invalid password: "+tc.Error))
			g.Expect(user).To(BeNil())
		})
	}

	t.Run("FailToFetchUser", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		email := "newuser@example.com"
		username := "New User"
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(email).Return(nil, errors.New("database error")).Times(1)

		service := auth.NewAuthService(repo, authmocks.NewMockPasswordHasherInterface(ctrl))

		user, err := service.Register(email, username, "password123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("failed to check existing user: database error"))
		g.Expect(user).To(BeNil())
	})

	t.Run("ExistingUser", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		email := "newuser@example.com"
		username := "New User"
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(email).Return(&models.User{}, nil).Times(1)

		service := auth.NewAuthService(repo, authmocks.NewMockPasswordHasherInterface(ctrl))

		user, err := service.Register(email, username, "password123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("user with this email already exists"))
		g.Expect(user).To(BeNil())
	})

	t.Run("FailedToHash", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		email := "newuser@example.com"
		username := "New User"
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(email).Return(nil, nil).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().Hash("password123").Return("", errors.New("hash error")).Times(1)

		service := auth.NewAuthService(repo, hasher)

		user, err := service.Register(email, username, "password123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("failed to hash password: hash error"))
		g.Expect(user).To(BeNil())
	})

	t.Run("FailedToCreateUser", func(t *testing.T) {
		g := NewWithT(t)
		ctrl := gomock.NewController(t)
		email := "newuser@example.com"
		hashedPassword := "hashed-password"
		username := "New User"
		repo := repomocks.NewMockUserRepository(ctrl)
		repo.EXPECT().GetByEmail(email).Return(nil, nil).Times(1)
		repo.EXPECT().Create(gomock.Any()).Return(errors.New("database error")).Times(1)

		hasher := authmocks.NewMockPasswordHasherInterface(ctrl)
		hasher.EXPECT().Hash("password123").Return(hashedPassword, nil).Times(1)
		service := auth.NewAuthService(repo, hasher)

		user, err := service.Register(email, username, "password123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err).To(MatchError("failed to create user: database error"))
		g.Expect(user).To(BeNil())
	})
}
