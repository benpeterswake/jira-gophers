package jira

import (
	"io"
	"net/http"
	"time"

	"gopkg.in/retry.v1"
)

type client struct {
	baseURL      string
	authURL      string
	scheme       string
	clientID     string
	clientSecret string
	redirectURI  string

	authService  AuthService
	issueService IssueService
}

type Client interface {
	NewRequest(method string, url string, body io.Reader) (*http.Request, error)
	SendRequest(req *http.Request) (*http.Response, error)
	GetBaseURL() string
	GetAuthUrl() string
	GetScheme() string
	GetClientID() string
	GetClientSecret() string
	GetRedirectURL() string
	GetAuthService() AuthService
	GetIssueService() IssueService
}

var attempts = retry.Regular{
	Total: 1 * time.Second,
	Delay: 250 * time.Millisecond,
	Min:   2,
}

func NewClient(domain string, authDomain string, scheme string, clientID string, clientSecret string, redirectURI string) Client {
	c := &client{
		baseURL:      domain,
		authURL:      authDomain,
		scheme:       scheme,
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
	}

	c.authService = &AuthImpl{c, "", ""}
	c.issueService = &IssueImpl{c}

	return c
}

func (c *client) GetBaseURL() string {
	return c.baseURL
}

func (c *client) GetAuthUrl() string {
	return c.authURL
}

func (c *client) GetScheme() string {
	return c.scheme
}

func (c *client) GetClientID() string {
	return c.clientID
}

func (c *client) GetClientSecret() string {
	return c.clientSecret
}

func (c *client) GetRedirectURL() string {
	return c.redirectURI
}

func (c *client) GetAuthService() AuthService {
	return c.authService
}

func (c *client) GetIssueService() IssueService {
	return c.issueService
}
