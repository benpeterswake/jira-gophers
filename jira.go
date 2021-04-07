package jira

import (
	"io"
	"net/http"
	"time"

	"gopkg.in/retry.v1"
)

type credentials struct {
	clientID     string
	clientSecret string
	redirectURI  string
}

type auth struct {
	accessToken  string
	refreshToken string
}
type request struct {
	baseURL string
	authURL string
	scheme  string
}

type client struct {
	credentials credentials
	request     request
	auth        auth
}

type Client interface {
	GetAccessTokenFromAuthorizationCode(code string) (*OAuthResponse, error)
	GetAccessTokenFromRefreshToken(refreshToken string) (*OAuthResponse, error)
	NewRequest(method string, url string, body io.Reader) (*http.Request, error)
	SendRequest(req *http.Request) (*http.Response, error)
	SetAuthData(accessToken string, refreshToken string)
	Search(jql string, options *SearchOptions) ([]Issue, error)
}

var attempts = retry.Regular{
	Total: 1 * time.Second,
	Delay: 250 * time.Millisecond,
	Min:   2,
}

func NewClient(domain string, clientID string, clientSecret string, redirectURI string) Client {
	return &client{
		credentials{
			clientID:     clientID,
			clientSecret: clientSecret,
			redirectURI:  redirectURI,
		},
		request{
			baseURL: domain,
			authURL: "auth.atlassian.com",
			scheme:  "https",
		},
		auth{},
	}
}
