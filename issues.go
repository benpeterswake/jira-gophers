package jira

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// Jira API docs: https://developer.atlassian.com/jiradev/jira-apis/jira-rest-apis/jira-rest-api-tutorials/jira-rest-api-example-query-issues
func (c *client) Search(jql string, options *SearchOptions) ([]Issue, error) {
	var v searchResult
	v.Issues = []Issue{}
	log.Println("[Search] Starting")
	for attempt := attempts.Start(nil); attempt.Next(); {
		log.Println("[Search] Starting Attempt:" + strconv.FormatInt(int64(attempt.Count()), 10))
		u := url.URL{
			Scheme: c.request.scheme,
			Host:   c.request.baseURL,
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

		req, err := c.NewRequest(method, u.String(), nil)
		if err != nil {
			return nil, err
		}

		resp, err := c.SendRequest(req)
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
			resp, err := c.GetAccessTokenFromRefreshToken(c.auth.refreshToken)
			if err != nil {
				return nil, err
			}
			c.SetAuthData(resp.AccessToken, c.auth.refreshToken)
			continue
		}
		break
	}
	log.Println("[Search] Ending")
	return v.Issues, nil
}
