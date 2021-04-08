package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type IssueImpl struct {
	client *client
}

type IssueService interface {
	Search(jql string, options *SearchOptions) ([]Issue, error)
	Update(key string, timeSpent string) error
}

// Jira API docs: https://developer.atlassian.com/jiradev/jira-apis/jira-rest-apis/jira-rest-api-tutorials/jira-rest-api-example-query-issues
func (i *IssueImpl) Search(jql string, options *SearchOptions) ([]Issue, error) {
	var v searchResult
	v.Issues = []Issue{}
	log.Println("[Search] Starting")
	for attempt := attempts.Start(nil); attempt.Next(); {
		log.Println("[Search] Starting Attempt:" + strconv.FormatInt(int64(attempt.Count()), 10))
		u := url.URL{
			Scheme: i.client.getScheme(),
			Host:   i.client.getBaseURL(),
			Path:   "rest/api/3/search",
		}
		uv := url.Values{}
		if jql != "" {
			uv.Add("jql", jql)
		}

		if options != nil {
			if options.StartAt != 0 {
				uv.Add("startAt", strconv.Itoa(options.StartAt))
			}
			if options.MaxResults != 0 {
				uv.Add("maxResults", strconv.Itoa(options.MaxResults))
			}
			if options.Expand != "" {
				uv.Add("expand", options.Expand)
			}
			if strings.Join(options.Fields, ",") != "" {
				uv.Add("fields", strings.Join(options.Fields, ","))
			}
			if options.ValidateQuery != "" {
				uv.Add("validateQuery", options.ValidateQuery)
			}
		}

		method := "GET"
		u.RawQuery = uv.Encode()

		req, err := i.client.newRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}

		resp, err := i.client.sendRequest(req)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(resp.Body).Decode(&v)
		if err != nil {
			return nil, err
		}

		// retry here
		if len(v.Issues) == 0 {
			// refresh token
			resp, err := i.client.GetAuthService().GetAccessTokenFromRefreshToken()
			if err != nil {
				return nil, err
			}
			i.client.GetAuthService().SetAccessToken(resp.AccessToken)
			continue
		}
		break
	}
	log.Println("[Search] Ending")
	return v.Issues, nil
}

func (i *IssueImpl) Update(key string, timeSpent string) error {

	log.Println("[Update] Starting")

	pathWithKey := fmt.Sprintf("rest/api/3/issue/%v/worklog", key)

	log.Println("path", pathWithKey)
	u := url.URL{
		Scheme: i.client.getScheme(),
		Host:   i.client.getBaseURL(),
		Path:   pathWithKey,
	}

	uv := url.Values{}
	method := "POST"
	u.RawQuery = uv.Encode()

	var workLog WorkLog

	workLog.TimeSpent = timeSpent

	requestBody, err := json.Marshal(&workLog)
	if err != nil {
		return err
	}

	req, err := i.client.newRequest(method, u.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	resp, err := i.client.sendRequest(req)
	if err != nil {
		return err
	}

	log.Println(resp)
	log.Println("[Update] Ending")
	return nil
}
