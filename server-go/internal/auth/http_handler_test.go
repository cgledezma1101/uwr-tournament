package auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/auth"
	"github.com/uwr-tournament/server-go/internal/auth/mocks"
	"go.uber.org/mock/gomock"
)

func TestHTTPHandler_LoginHandler_SuccessfulLogin(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	hasher := mocks.NewMockPasswordHasherInterface(ctrl)
	authService := mocks.NewMockAuthService(ctrl)
	password := "testPassword123"
	user := &auth.User{
		ID:       1,
		Email:    "test@example.com",
		Name:     "Test User",
		Provider: "provider",
	}
	authService.EXPECT().AuthenticateWithPassword("test@example.com", password).Return(user, nil).Times(1)
	handler := auth.NewHTTPHandler(authService, hasher, mocks.NewMockOAuthProvider(ctrl), mocks.NewMockOAuthProvider(ctrl))
	reqBody := auth.LoginRequest{
		Email:    "test@example.com",
		Password: password,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v2/auth/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.LoginHandler(w, req)

	g.Expect(w.Code).To(Equal(http.StatusOK))

	var response auth.AuthResponse
	g.Expect(json.NewDecoder(w.Body).Decode(&response)).To(Succeed())
	g.Expect(response.ID).To(Equal(user.ID))
	g.Expect(response.Email).To(Equal(user.Email))
	g.Expect(response.Name).To(Equal(user.Name))
	g.Expect(response.Provider).To(Equal(user.Provider))
}

func TestHTTPHandler_LoginHandler_AuthorizerError(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	hasher := mocks.NewMockPasswordHasherInterface(ctrl)
	authService := mocks.NewMockAuthService(ctrl)
	authService.EXPECT().AuthenticateWithPassword(gomock.Any(), gomock.Any()).Return(nil, errors.New("some authN error")).Times(1)
	handler := auth.NewHTTPHandler(authService, hasher, mocks.NewMockOAuthProvider(ctrl), mocks.NewMockOAuthProvider(ctrl))

	// Create request with wrong password
	reqBody := auth.LoginRequest{
		Email:    "test@example.com",
		Password: "wrongPassword",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v2/auth/login", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.LoginHandler(w, req)

	g.Expect(w.Code).To(Equal(http.StatusUnauthorized))
}

func TestHTTPHandler_RegisterHandler_SuccessfulRegistration(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	hasher := mocks.NewMockPasswordHasherInterface(ctrl)
	authService := mocks.NewMockAuthService(ctrl)
	email := "newuser@example.com"
	name := "New User"
	password := "securePassword123"
	authUser := &auth.User{
		ID:       2,
		Email:    email,
		Name:     name,
		Provider: "provider",
	}
	authService.EXPECT().Register(email, name, password).Return(authUser, nil).Times(1)
	handler := auth.NewHTTPHandler(authService, hasher, mocks.NewMockOAuthProvider(ctrl), mocks.NewMockOAuthProvider(ctrl))

	reqBody := auth.RegisterRequest{
		Email:    email,
		Name:     name,
		Password: password,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v2/auth/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.RegisterHandler(w, req)

	g.Expect(w.Code).To(Equal(http.StatusCreated))

	var response auth.AuthResponse
	g.Expect(json.NewDecoder(w.Body).Decode(&response)).To(Succeed())
	g.Expect(response.Email).To(Equal(authUser.Email))
	g.Expect(response.Name).To(Equal(authUser.Name))
	g.Expect(response.Provider).To(Equal(authUser.Provider))
	g.Expect(response.ID).To(Equal(authUser.ID))
}

func TestHTTPHandler_RegisterHandler_RegistrationError(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	hasher := mocks.NewMockPasswordHasherInterface(ctrl)
	authService := mocks.NewMockAuthService(ctrl)
	authService.EXPECT().Register(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("registration error")).Times(1)
	handler := auth.NewHTTPHandler(authService, hasher, mocks.NewMockOAuthProvider(ctrl), mocks.NewMockOAuthProvider(ctrl))

	reqBody := auth.RegisterRequest{
		Email:    "duplicate@example.com",
		Name:     "User Two",
		Password: "password456",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v2/auth/register", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.RegisterHandler(w, req)

	g.Expect(w.Code).To(Equal(http.StatusBadRequest))
}

func TestHTTPHandler_GoogleOAuthHandler_RedirectsToGoogle(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	googleProvider := mocks.NewMockOAuthProvider(ctrl)
	googleProvider.EXPECT().GetAuthorizationURL(gomock.Any()).Return("https://expected.redirect.com").Times(1)
	handler := auth.NewHTTPHandler(
		mocks.NewMockAuthService(ctrl),
		mocks.NewMockPasswordHasherInterface(ctrl),
		googleProvider,
		mocks.NewMockOAuthProvider(ctrl))

	req := httptest.NewRequest(http.MethodGet, "/oauth/google?state=test_state", nil)
	w := httptest.NewRecorder()

	handler.GoogleOAuthHandler(w, req)

	g.Expect(w.Code).To(Equal(http.StatusTemporaryRedirect))

	locationHeader := w.Header().Get("Location")
	g.Expect(locationHeader).To(Equal("https://expected.redirect.com"))
}

func TestHTTPHandler_GoogleOAuthCallbackHandler_SuccessfulCallback(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	googleProvider := mocks.NewMockOAuthProvider(ctrl)
	googleProvider.EXPECT().ExchangeCodeForToken("auth_code").Return("valid_access_token", nil).Times(1)

	oauthData := &auth.OAuthData{
		Email: "user@example.com",
		Name:  "Test User",
	}
	googleProvider.EXPECT().GetUserInfo("valid_access_token").Return(oauthData, nil).Times(1)

	user := &auth.User{
		ID:       3,
		Email:    "created@example.com",
		Name:     "Created User",
		Provider: "google",
	}
	authService := mocks.NewMockAuthService(ctrl)
	authService.EXPECT().GetOrCreateFromOAuth(oauthData).Return(user, nil).Times(1)

	handler := auth.NewHTTPHandler(
		authService,
		mocks.NewMockPasswordHasherInterface(ctrl),
		googleProvider,
		mocks.NewMockOAuthProvider(ctrl))

	req := httptest.NewRequest(http.MethodGet, "/oauth/google/callback?code=auth_code&state=test_state", nil)
	w := httptest.NewRecorder()

	handler.GoogleOAuthCallbackHandler(w, req)
	g.Expect(w.Code).To(Equal(http.StatusOK))

	var response auth.AuthResponse
	g.Expect(json.Unmarshal(w.Body.Bytes(), &response)).To(Succeed())

	g.Expect(response.Email).To(Equal(user.Email))
	g.Expect(response.Name).To(Equal(user.Name))
	g.Expect(response.Provider).To(Equal(user.Provider))
	g.Expect(response.ID).To(Equal(user.ID))
}

func TestHTTPHandler_FacebookOAuthHandler_RedirectsToFacebook(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	facebookProvider := mocks.NewMockOAuthProvider(ctrl)
	facebookProvider.EXPECT().GetAuthorizationURL(gomock.Any()).Return("https://expected.redirect.com").Times(1)
	handler := auth.NewHTTPHandler(
		mocks.NewMockAuthService(ctrl),
		mocks.NewMockPasswordHasherInterface(ctrl),
		mocks.NewMockOAuthProvider(ctrl),
		facebookProvider)

	req := httptest.NewRequest(http.MethodGet, "/oauth/facebook?state=test_state", nil)
	w := httptest.NewRecorder()

	handler.FacebookOAuthHandler(w, req)

	g.Expect(w.Code).To(Equal(http.StatusTemporaryRedirect))

	locationHeader := w.Header().Get("Location")
	g.Expect(locationHeader).To(Equal("https://expected.redirect.com"))
}

func TestHTTPHandler_FacebookOAuthCallbackHandler_SuccessfulCallback(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	facebookProvider := mocks.NewMockOAuthProvider(ctrl)
	facebookProvider.EXPECT().ExchangeCodeForToken("auth_code").Return("valid_access_token", nil).Times(1)

	oauthData := &auth.OAuthData{
		Email: "user@example.com",
		Name:  "Test User",
	}
	facebookProvider.EXPECT().GetUserInfo("valid_access_token").Return(oauthData, nil).Times(1)

	user := &auth.User{
		ID:       3,
		Email:    "created@example.com",
		Name:     "Created User",
		Provider: "facebook",
	}
	authService := mocks.NewMockAuthService(ctrl)
	authService.EXPECT().GetOrCreateFromOAuth(oauthData).Return(user, nil).Times(1)

	handler := auth.NewHTTPHandler(
		authService,
		mocks.NewMockPasswordHasherInterface(ctrl),
		mocks.NewMockOAuthProvider(ctrl),
		facebookProvider)

	req := httptest.NewRequest(http.MethodGet, "/oauth/facebook/callback?code=auth_code&state=test_state", nil)
	w := httptest.NewRecorder()

	handler.FacebookOAuthCallbackHandler(w, req)
	g.Expect(w.Code).To(Equal(http.StatusOK))

	var response auth.AuthResponse
	g.Expect(json.Unmarshal(w.Body.Bytes(), &response)).To(Succeed())

	g.Expect(response.Email).To(Equal(user.Email))
	g.Expect(response.Name).To(Equal(user.Name))
	g.Expect(response.Provider).To(Equal(user.Provider))
	g.Expect(response.ID).To(Equal(user.ID))
}

func TestHTTPHandler_RequestBodyParsing_InvalidJSON(t *testing.T) {
	g := NewWithT(t)
	ctrl := gomock.NewController(t)
	handler := auth.NewHTTPHandler(mocks.NewMockAuthService(ctrl),
		mocks.NewMockPasswordHasherInterface(ctrl),
		mocks.NewMockOAuthProvider(ctrl),
		mocks.NewMockOAuthProvider(ctrl))

	req := httptest.NewRequest(http.MethodPost, "/api/v2/auth/login", io.NopCloser(bytes.NewReader([]byte("invalid json"))))
	w := httptest.NewRecorder()	

	handler.LoginHandler(w, req)

	g.Expect(w.Code).To(Equal(http.StatusBadRequest))
}
