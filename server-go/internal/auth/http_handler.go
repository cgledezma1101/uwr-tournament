package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/uwr-tournament/server-go/internal/responses"
)

// LoginRequest represents a login request body
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest represents a registration request body
type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// AuthResponse represents a successful authentication response
type AuthResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// OAuthCallbackRequest represents the OAuth callback parameters
type OAuthCallbackRequest struct {
	Code  string
	State string
}

// HTTPHandler provides HTTP handlers for authentication
type HTTPHandler struct {
	authService      AuthService
	passwordHasher   PasswordHasherInterface
	googleProvider   OAuthProvider
	facebookProvider OAuthProvider
}

// NewHTTPHandler creates a new HTTP handler
func NewHTTPHandler(authService AuthService, passwordHasher PasswordHasherInterface, googleProvider, facebookProvider OAuthProvider) *HTTPHandler {
	return &HTTPHandler{
		authService:      authService,
		passwordHasher:   passwordHasher,
		googleProvider:   googleProvider,
		facebookProvider: facebookProvider,
	}
}

// LoginHandler handles username/password login
// POST /api/v2/auth/login
func (h *HTTPHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.authService.AuthenticateWithPassword(req.Email, req.Password)
	if err != nil {
		responses.SendError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	responses.SendJSON(w, h.userToResponse(user), http.StatusOK)
}

// RegisterHandler handles user registration
// POST /api/v2/auth/register
func (h *HTTPHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.SendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.authService.Register(req.Email, req.Name, req.Password)
	if err != nil {
		responses.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.SendJSON(w, h.userToResponse(user), http.StatusCreated)
}

// OAuthInitiatorGenerator creates an OAuth initiation handler for a given provider
func (h *HTTPHandler) OAuthInitiatorGenerator(provider OAuthProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		if state == "" {
			state = h.passwordHasher.GenerateToken(32)
		}

		authURL := provider.GetAuthorizationURL(state)
		http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
	}
}

// OAuthCallbackGenerator creates an OAuth callback handler for a given provider
func (h *HTTPHandler) OAuthCallbackGenerator(provider OAuthProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		if code == "" {
			errorMsg := r.URL.Query().Get("error")
			responses.SendError(w, fmt.Sprintf("OAuth error: %s", errorMsg), http.StatusUnauthorized)
			return
		}

		token, err := provider.ExchangeCodeForToken(code)
		if err != nil {
			responses.SendError(w, "Failed to exchange code for token", http.StatusInternalServerError)
			return
		}

		oauthData, err := provider.GetUserInfo(token)
		if err != nil {
			responses.SendError(w, "Failed to get user info", http.StatusInternalServerError)
			return
		}

		user, err := h.authService.GetOrCreateFromOAuth(oauthData)
		if err != nil {
			responses.SendError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responses.SendJSON(w, h.userToResponse(user), http.StatusOK)
	}
}

// GoogleOAuthHandler initiates Google OAuth flow
// GET /oauth/google
func (h *HTTPHandler) GoogleOAuthHandler(w http.ResponseWriter, r *http.Request) {
	h.OAuthInitiatorGenerator(h.googleProvider)(w, r)
}

// GoogleOAuthCallbackHandler handles Google OAuth callback
// GET /oauth/google/callback
func (h *HTTPHandler) GoogleOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	h.OAuthCallbackGenerator(h.googleProvider)(w, r)
}

// FacebookOAuthHandler initiates Facebook OAuth flow
// GET /oauth/facebook
func (h *HTTPHandler) FacebookOAuthHandler(w http.ResponseWriter, r *http.Request) {
	h.OAuthInitiatorGenerator(h.facebookProvider)(w, r)
}

// FacebookOAuthCallbackHandler handles Facebook OAuth callback
// GET /oauth/facebook/callback
func (h *HTTPHandler) FacebookOAuthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	h.OAuthCallbackGenerator(h.facebookProvider)(w, r)
}

// Helper functions

func (h *HTTPHandler) userToResponse(user *User) *AuthResponse {
	if user == nil {
		return nil
	}
	return &AuthResponse{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Provider: user.Provider,
	}
}
