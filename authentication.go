package jira

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type AuthImpl struct {
	client       *client
	accessToken  string
	refreshToken string
}

type AuthService interface {
	GetRefreshToken() string
	GetAccessToken() string
	SetAccessToken(accessToken string)
	SetRefreshToken(refreshToken string)
	GetAccessTokenFromAuthorizationCode(code string) (*OAuthResponse, error)
	GetAccessTokenFromRefreshToken() (*OAuthResponse, error)
}

type OAuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type OAuthRefreshRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

type OAuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func (c *AuthImpl) GetRefreshToken() string {
	return c.refreshToken
}

func (c *AuthImpl) GetAccessToken() string {
	return c.accessToken
}

func (c *AuthImpl) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}

func (c *AuthImpl) SetRefreshToken(refreshToken string) {
	c.refreshToken = refreshToken
}

func (a *AuthImpl) GetAccessTokenFromAuthorizationCode(code string) (*OAuthResponse, error) {
	u := url.URL{
		Scheme: a.client.getScheme(),
		Host:   a.client.getAuthUrl(),
		Path:   "oauth/token",
	}

	method := "POST"

	payload := OAuthRequest{
		GrantType:    "authorization_code",
		ClientID:     a.client.getClientID(),
		ClientSecret: a.client.getClientSecret(),
		Code:         code,
		RedirectURI:  a.client.getRedirectURL(),
	}

	requestBody, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	log.Println(u.String())

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bytesResp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("[GetAccessTokenFromAuthorizationCode] Error reading response body" + err.Error())
			return nil, err
		}
		log.Println("[GetAccessTokenFromAuthorizationCode] Error body", string(bytesResp))
		log.Println("[GetAccessTokenFromAuthorizationCode] Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))

		return nil, errors.New("[GetAccessTokenFromAuthorizationCode] Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))
	}

	var resp OAuthResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (a *AuthImpl) GetAccessTokenFromRefreshToken() (*OAuthResponse, error) {
	u := url.URL{
		Scheme: a.client.getScheme(),
		Host:   a.client.getAuthUrl(),
		Path:   "oauth/token",
	}

	method := "POST"

	payload := OAuthRefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     a.client.getClientID(),
		ClientSecret: a.client.getClientSecret(),
		RefreshToken: a.refreshToken,
	}

	requestBody, err := json.Marshal(&payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		bytesResp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("[GetAccessTokenFromRefreshToken] Error reading response body" + err.Error())
			return nil, err
		}
		log.Println("[GetAccessTokenFromRefreshToken] Error body", string(bytesResp))
		log.Println("[GetAccessTokenFromRefreshToken] Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))
		return nil, errors.New("[GetAccessTokenFromRefreshToken] Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))
	}

	var resp OAuthResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}
