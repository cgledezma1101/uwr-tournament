package auth_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/uwr-tournament/server-go/internal/auth"
)

func TestGoogleOAuthProvider(t *testing.T) {
	t.Run("GetAuthorizationURL", func(t *testing.T) {
		g := NewWithT(t)
		provider := auth.NewGoogleOAuthProvider("test_client_id", "test_secret", "http://localhost:3000/oauth/google/callback", "https://auth.base.url", "", "")
		state := "test_state"

		actualUrl := provider.GetAuthorizationURL(state)
		g.Expect(actualUrl).To(And(
			ContainSubstring("https://auth.base.url/o/oauth2/v2/auth"),
			ContainSubstring("client_id=test_client_id"),
			ContainSubstring("state=test_state"),
			ContainSubstring("redirect_uri="+url.QueryEscape("http://localhost:3000/oauth/google/callback")),
			ContainSubstring("response_type=code"),
			ContainSubstring("scope=openid+email+profile")))
	})

	t.Run("ExchangeCodeForToken", func(t *testing.T) {
		g := NewWithT(t)
		authCode := "auth_code_123"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			g.Expect(r.URL.Path).To(Equal("/token"))
			g.Expect(r.Method).To(Equal(http.MethodPost))
			g.Expect(r.Header.Get("Content-Type")).To(Equal("application/x-www-form-urlencoded"))
			payload, err := io.ReadAll(r.Body)

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(string(payload)).To(And(
				ContainSubstring("code="+authCode),
				ContainSubstring("client_id=client_id"),
				ContainSubstring("client_secret=client_secret"),
				ContainSubstring("redirect_uri=http://reditect.uri"),
				ContainSubstring("grant_type=authorization_code")))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"test_access_token","token_type":"Bearer","expires_in":3600}`))
		}))
		provider := auth.NewGoogleOAuthProvider("client_id", "client_secret", "http://reditect.uri", "", testServer.URL, "")

		token, err := provider.ExchangeCodeForToken(authCode)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(token).To(Equal("test_access_token"))
	})

	t.Run("ExchangeCodeForToken_InvalidResponse", func(t *testing.T) {
		g := NewWithT(t)
		authCode := "auth_code_123"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("not_a_token"))
		}))
		provider := auth.NewGoogleOAuthProvider("client_id", "client_secret", "http://reditect.uri", "", testServer.URL, "")

		_, err := provider.ExchangeCodeForToken(authCode)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("failed to parse token response"))
	})

	t.Run("ExchangeCodeForToken_NonOkResponse", func(t *testing.T) {
		g := NewWithT(t)
		authCode := "auth_code_123"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("not_a_token"))
		}))
		provider := auth.NewGoogleOAuthProvider("client_id", "client_secret", "http://reditect.uri", "", testServer.URL, "")

		_, err := provider.ExchangeCodeForToken(authCode)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("token exchange failed with status 500"))
	})

	t.Run("ExchangeCodeForToken_EmptyCode", func(t *testing.T) {
		g := NewWithT(t)
		provider := auth.NewGoogleOAuthProvider("", "", "", "", "", "")

		_, err := provider.ExchangeCodeForToken("")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("authorization code cannot be empty"))

		if err == nil {
			t.Error("ExchangeCodeForToken() should return error for empty code")
		}
	})

	t.Run("GetUserInfo", func(t *testing.T) {
		g := NewWithT(t)
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			g.Expect(r.URL.Path).To(Equal("/oauth2/v1/userinfo"))
			g.Expect(r.Method).To(Equal(http.MethodGet))
			g.Expect(r.URL.Query().Get("alt")).To(Equal("json"))
			g.Expect(r.Header.Get("Authorization")).To(Equal("Bearer access_token_123"))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"email":"test@example.com","name":"Test User","id":"123"}`))
		}))

		provider := auth.NewGoogleOAuthProvider("", "", "", "", "", testServer.URL)

		userInfo, err := provider.GetUserInfo("access_token_123")
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(userInfo.Name).To(Equal("Test User"))
		g.Expect(userInfo.Email).To(Equal("test@example.com"))
		g.Expect(userInfo.UID).To(Equal("123"))
		g.Expect(userInfo.Provider).To(Equal("google"))
	})

	t.Run("GetUserInfo_InvalidBody", func(t *testing.T) {
		g := NewWithT(t)
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("not_json"))
		}))

		provider := auth.NewGoogleOAuthProvider("", "", "", "", "", testServer.URL)

		userInfo, err := provider.GetUserInfo("access_token_123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("failed to parse user info"))
		g.Expect(userInfo).To(BeNil())
	})

	t.Run("GetUserInfo_NonSuccess", func(t *testing.T) {
		g := NewWithT(t)
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("not_json"))
		}))

		provider := auth.NewGoogleOAuthProvider("", "", "", "", "", testServer.URL)

		userInfo, err := provider.GetUserInfo("access_token_123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("user info retrieval failed with status 500"))
		g.Expect(userInfo).To(BeNil())
	})
}

func TestFacebookOAuthProvider(t *testing.T) {
	t.Run("GetAuthorizationURL", func(t *testing.T) {
		g := NewWithT(t)
		provider := auth.NewFacebookOAuthProvider("test_client_id", "test_secret", "http://localhost:3000/oauth/facebook/callback", "https://auth.base.url", "", "")
		state := "test_state"

		actualUrl := provider.GetAuthorizationURL(state)
		g.Expect(actualUrl).To(And(
			ContainSubstring("https://auth.base.url/v12.0/dialog/oauth"),
			ContainSubstring("client_id=test_client_id"),
			ContainSubstring("state=test_state"),
			ContainSubstring("redirect_uri="+url.QueryEscape("http://localhost:3000/oauth/facebook/callback")),
			ContainSubstring("scope=email")))
	})

	t.Run("ExchangeCodeForToken", func(t *testing.T) {
		g := NewWithT(t)
		authCode := "auth_code_123"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			g.Expect(r.URL.Path).To(Equal("/v12.0/oauth/access_token"))
			g.Expect(r.Method).To(Equal(http.MethodPost))
			g.Expect(r.Header.Get("Content-Type")).To(Equal("application/x-www-form-urlencoded"))
			payload, err := io.ReadAll(r.Body)

			g.Expect(err).NotTo(HaveOccurred())
			g.Expect(string(payload)).To(And(
				ContainSubstring("code="+authCode),
				ContainSubstring("client_id=client_id"),
				ContainSubstring("client_secret=client_secret"),
				ContainSubstring("redirect_uri=http://reditect.uri")))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"test_access_token","token_type":"Bearer","expires_in":3600}`))
		}))
		provider := auth.NewFacebookOAuthProvider("client_id", "client_secret", "http://reditect.uri", "", testServer.URL, "")

		token, err := provider.ExchangeCodeForToken(authCode)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(token).To(Equal("test_access_token"))
	})

	t.Run("ExchangeCodeForToken_InvalidResponse", func(t *testing.T) {
		g := NewWithT(t)
		authCode := "auth_code_123"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("not_a_token"))
		}))
		provider := auth.NewFacebookOAuthProvider("client_id", "client_secret", "http://reditect.uri", "", testServer.URL, "")

		_, err := provider.ExchangeCodeForToken(authCode)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("failed to parse token response"))
	})

	t.Run("ExchangeCodeForToken_NonOkResponse", func(t *testing.T) {
		g := NewWithT(t)
		authCode := "auth_code_123"
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("not_a_token"))
		}))
		provider := auth.NewFacebookOAuthProvider("client_id", "client_secret", "http://reditect.uri", "", testServer.URL, "")

		_, err := provider.ExchangeCodeForToken(authCode)
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("token exchange failed with status 500"))
	})

	t.Run("ExchangeCodeForToken_EmptyCode", func(t *testing.T) {
		g := NewWithT(t)
		provider := auth.NewFacebookOAuthProvider("", "", "", "", "", "")

		_, err := provider.ExchangeCodeForToken("")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("authorization code cannot be empty"))

		if err == nil {
			t.Error("ExchangeCodeForToken() should return error for empty code")
		}
	})

	t.Run("GetUserInfo", func(t *testing.T) {
		g := NewWithT(t)
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			g.Expect(r.URL.Path).To(Equal("/me"))
			g.Expect(r.Method).To(Equal(http.MethodGet))
			g.Expect(r.URL.Query().Get("fields")).To(Equal("id,email,name"))
			g.Expect(r.URL.Query().Get("access_token")).To(Equal("access_token_123"))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"email":"test@example.com","name":"Test User","id":"123"}`))
		}))

		provider := auth.NewFacebookOAuthProvider("", "", "", "", "", testServer.URL)

		userInfo, err := provider.GetUserInfo("access_token_123")
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(userInfo.Name).To(Equal("Test User"))
		g.Expect(userInfo.Email).To(Equal("test@example.com"))
		g.Expect(userInfo.UID).To(Equal("123"))
		g.Expect(userInfo.Provider).To(Equal("facebook"))
	})

	t.Run("GetUserInfo_InvalidBody", func(t *testing.T) {
		g := NewWithT(t)
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("not_json"))
		}))

		provider := auth.NewFacebookOAuthProvider("", "", "", "", "", testServer.URL)

		userInfo, err := provider.GetUserInfo("access_token_123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("failed to parse user info"))
		g.Expect(userInfo).To(BeNil())
	})

	t.Run("GetUserInfo_NonSuccess", func(t *testing.T) {
		g := NewWithT(t)
		testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("not_json"))
		}))

		provider := auth.NewFacebookOAuthProvider("", "", "", "", "", testServer.URL)

		userInfo, err := provider.GetUserInfo("access_token_123")
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(ContainSubstring("user info retrieval failed with status 500"))
		g.Expect(userInfo).To(BeNil())
	})
}
