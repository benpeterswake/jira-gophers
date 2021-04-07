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

func (c *client) SetAuthData(accessToken string, refreshToken string) {
	c.auth.accessToken = accessToken
	c.auth.refreshToken = refreshToken
}

func (c *client) GetAccessTokenFromAuthorizationCode(code string) (*OAuthResponse, error) {
	u := url.URL{
		Scheme: c.request.scheme,
		Host:   c.request.authURL,
		Path:   "oauth/token",
	}

	method := "POST"

	payload := OAuthRequest{
		GrantType:    "authorization_code",
		ClientID:     c.credentials.clientID,
		ClientSecret: c.credentials.clientSecret,
		Code:         code,
		RedirectURI:  c.credentials.redirectURI,
	}

	requestBody, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

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
			log.Println("Error reading response body" + err.Error())
			return nil, err
		}
		log.Println("Error body", string(bytesResp))
		log.Println("Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))

		return nil, errors.New("Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))
	}

	var resp OAuthResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &resp, nil
}

func (c *client) GetAccessTokenFromRefreshToken(refreshToken string) (*OAuthResponse, error) {
	u := url.URL{
		Scheme: c.request.scheme,
		Host:   c.request.authURL,
		Path:   "oauth/token",
	}

	method := "POST"

	payload := OAuthRefreshRequest{
		GrantType:    "refresh_token",
		ClientID:     c.credentials.clientID,
		ClientSecret: c.credentials.clientSecret,
		RefreshToken: refreshToken,
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
			log.Println("Error reading response body" + err.Error())
			return nil, err
		}
		log.Println("Error body", string(bytesResp))
		log.Println("Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))
		return nil, errors.New("Error calling jira auth api. Wanted 200 but got code " + strconv.FormatInt(int64(res.StatusCode), 10))
	}

	var resp OAuthResponse
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &resp, nil
}
