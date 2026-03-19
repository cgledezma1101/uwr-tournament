package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GoogleOAuthProvider handles Google OAuth authentication
type GoogleOAuthProvider struct {
	clientID             string
	clientSecret         string
	redirectURI          string
	authorizationBaseUrl string
	tokenBaseUrl         string
	userInfoBaseUrl      string
	httpClient           *http.Client
}

// NewGoogleOAuthProvider creates a new Google OAuth provider
func NewGoogleOAuthProvider(clientID, clientSecret, redirectURI, authorizationBaseUrl, tokenBaseUrl, userInfoBaseUrl string) OAuthProvider {
	return &GoogleOAuthProvider{
		clientID:             clientID,
		clientSecret:         clientSecret,
		redirectURI:          redirectURI,
		authorizationBaseUrl: authorizationBaseUrl,
		tokenBaseUrl:         tokenBaseUrl,
		userInfoBaseUrl:      userInfoBaseUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetAuthorizationURL returns the Google OAuth authorization URL
func (g *GoogleOAuthProvider) GetAuthorizationURL(state string) string {
	return fmt.Sprintf(
		"%s/o/oauth2/v2/auth?client_id=%s&redirect_uri=%s&response_type=code&scope=openid+email+profile&state=%s",
		g.authorizationBaseUrl,
		url.QueryEscape(g.clientID),
		url.QueryEscape(g.redirectURI),
		url.QueryEscape(state),
	)
}

// ExchangeCodeForToken exchanges an authorization code for an access token
func (g *GoogleOAuthProvider) ExchangeCodeForToken(code string) (string, error) {
	if code == "" {
		return "", errors.New("authorization code cannot be empty")
	}

	tokenURL := fmt.Sprintf("%s/token", g.tokenBaseUrl)
	payload := fmt.Sprintf(
		"code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code",
		code,
		g.clientID,
		g.clientSecret,
		g.redirectURI,
	)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to exchange token: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token exchange failed with status %d", resp.StatusCode)
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}

	return tokenResp.AccessToken, nil
}

// GetUserInfo retrieves user information from Google using the access token
func (g *GoogleOAuthProvider) GetUserInfo(accessToken string) (*OAuthData, error) {
	if accessToken == "" {
		return nil, errors.New("access token cannot be empty")
	}

	userInfoURL := fmt.Sprintf("%s/oauth2/v1/userinfo?alt=json", g.userInfoBaseUrl)
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create user info request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user info retrieval failed with status %d", resp.StatusCode)
	}

	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		ID    string `json:"id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse user info: %w", err)
	}

	return &OAuthData{
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Provider: "google",
		UID:      userInfo.ID,
	}, nil
}

// FacebookOAuthProvider handles Facebook OAuth authentication
type FacebookOAuthProvider struct {
	clientID             string
	clientSecret         string
	redirectURI          string
	authorizationBaseUrl string
	tokenBaseUrl         string
	userInfoBaseUrl      string
	httpClient           *http.Client
}

// NewFacebookOAuthProvider creates a new Facebook OAuth provider
func NewFacebookOAuthProvider(clientID, clientSecret, redirectURI, authorizationBaseUrl, tokenBaseUrl, userInfoBaseUrl string) OAuthProvider {
	return &FacebookOAuthProvider{
		clientID:             clientID,
		clientSecret:         clientSecret,
		redirectURI:          redirectURI,
		authorizationBaseUrl: authorizationBaseUrl,
		tokenBaseUrl:         tokenBaseUrl,
		userInfoBaseUrl:      userInfoBaseUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetAuthorizationURL returns the Facebook OAuth authorization URL
func (f *FacebookOAuthProvider) GetAuthorizationURL(state string) string {
	return fmt.Sprintf(
		"%s/v12.0/dialog/oauth?client_id=%s&redirect_uri=%s&scope=email&state=%s",
		f.authorizationBaseUrl,
		url.QueryEscape(f.clientID),
		url.QueryEscape(f.redirectURI),
		url.QueryEscape(state),
	)
}

// ExchangeCodeForToken exchanges an authorization code for an access token
func (f *FacebookOAuthProvider) ExchangeCodeForToken(code string) (string, error) {
	if code == "" {
		return "", errors.New("authorization code cannot be empty")
	}

	tokenURL := fmt.Sprintf("%s/v12.0/oauth/access_token", f.tokenBaseUrl)
	payload := fmt.Sprintf(
		"client_id=%s&client_secret=%s&redirect_uri=%s&code=%s",
		f.clientID,
		f.clientSecret,
		f.redirectURI,
		code,
	)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to exchange token: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token exchange failed with status %d", resp.StatusCode)
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}

	return tokenResp.AccessToken, nil
}

// GetUserInfo retrieves user information from Facebook using the access token
func (f *FacebookOAuthProvider) GetUserInfo(accessToken string) (*OAuthData, error) {
	if accessToken == "" {
		return nil, errors.New("access token cannot be empty")
	}

	userInfoURL := fmt.Sprintf("%s/me?fields=id,email,name&access_token=%s", f.userInfoBaseUrl, accessToken)
	req, err := http.NewRequest("GET", userInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create user info request: %w", err)
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user info retrieval failed with status %d", resp.StatusCode)
	}

	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		ID    string `json:"id"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse user info: %w", err)
	}

	return &OAuthData{
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Provider: "facebook",
		UID:      userInfo.ID,
	}, nil
}
