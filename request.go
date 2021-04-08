package jira

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func (c *client) newRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req, nil
}

func (c *client) sendRequest(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	log.Println("[SendRequest] Started")
	for attempt := attempts.Start(nil); attempt.Next(); {
		log.Println("[SendRequest] Starting Attempt:" + strconv.FormatInt(int64(attempt.Count()), 10))

		if c.GetAuthService().GetAccessToken() == "" {
			// refresh token
			log.Println("Token here:", c.GetAuthService().GetRefreshToken())
			authResp, err := c.GetAuthService().GetAccessTokenFromRefreshToken()
			if err != nil {
				log.Println(err)
				return nil, err
			}
			c.GetAuthService().SetAccessToken(authResp.AccessToken)
		}

		req.Header.Set("Authorization", "Bearer "+c.GetAuthService().GetAccessToken())

		httpClient := http.Client{}
		var err error
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println("Error sending request" + err.Error())
			return nil, err
		}

		if resp.StatusCode != 200 {
			bytesResp, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Error reading response body" + err.Error())
				return nil, err
			}

			log.Println("Error body", string(bytesResp))
			log.Println("Error calling jira api. Wanted 200 but got code " + strconv.FormatInt(int64(resp.StatusCode), 10))
			// if error is unauthorized retry here
			if resp.StatusCode == 401 {
				c.GetAuthService().SetAccessToken("")
				continue
			}
			return nil, errors.New("Error calling jira api. Wanted 200 but got code " + strconv.FormatInt(int64(resp.StatusCode), 10))
		}
		break
	}
	log.Println("[SendRequest] Ended")
	return resp, nil
}
