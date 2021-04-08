package jira

import (
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

func (c *client) getBaseURL() string {
	return c.baseURL
}

func (c *client) getAuthUrl() string {
	return c.authURL
}

func (c *client) getScheme() string {
	return c.scheme
}

func (c *client) getClientID() string {
	return c.clientID
}

func (c *client) getClientSecret() string {
	return c.clientSecret
}

func (c *client) getRedirectURL() string {
	return c.redirectURI
}

func (c *client) GetAuthService() AuthService {
	return c.authService
}

func (c *client) GetIssueService() IssueService {
	return c.issueService
}
