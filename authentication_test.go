package jira

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestClient_GetAccessTokenFromRefreshToken(t *testing.T) {
	setup()
	defer teardown()

	refreshToken := "test"

	expected := OAuthResponse{
		AccessToken:  "test token",
		RefreshToken: refreshToken,
		Scope:        "scope",
		ExpiresIn:    3600,
		TokenType:    "test type",
	}

	testMux.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		resp := expected
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			t.Fatal(err)
		}
		return
	})

	actual, err := testClient.GetAuthService().GetAccessTokenFromRefreshToken()
	if err != nil {
		t.Fatal(err)
	}

	if actual.AccessToken != expected.AccessToken {
		t.Fatal("Wanted", expected.AccessToken, "but got", actual.AccessToken)
	}

	if actual.RefreshToken != expected.RefreshToken {
		t.Fatal("Wanted", expected.RefreshToken, "but got", actual.RefreshToken)
	}

	if actual.Scope != expected.Scope {
		t.Fatal("Wanted", expected.Scope, "but got", actual.Scope)
	}

	if actual.ExpiresIn != expected.ExpiresIn {
		t.Fatal("Wanted", expected.ExpiresIn, "but got", actual.ExpiresIn)
	}

	if actual.TokenType != expected.TokenType {
		t.Fatal("Wanted", expected.TokenType, "but got", actual.TokenType)
	}
}

func TestClient_GetAccessTokenFromAuthorizationCode(t *testing.T) {
	setup()
	defer teardown()

	code := "test"

	expected := OAuthResponse{
		AccessToken:  code,
		RefreshToken: "refresh token",
		Scope:        "scope",
		ExpiresIn:    3600,
		TokenType:    "test type",
	}

	testMux.HandleFunc("/oauth/token", func(w http.ResponseWriter, r *http.Request) {
		resp := expected
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(&resp)
		if err != nil {
			t.Fatal(err)
		}
		return
	})

	actual, err := testClient.GetAuthService().GetAccessTokenFromAuthorizationCode(code)
	if err != nil {
		t.Fatal(err)
	}

	if actual.AccessToken != expected.AccessToken {
		t.Fatal("Wanted", expected.AccessToken, "but got", actual.AccessToken)
	}

	if actual.RefreshToken != expected.RefreshToken {
		t.Fatal("Wanted", expected.RefreshToken, "but got", actual.RefreshToken)
	}

	if actual.Scope != expected.Scope {
		t.Fatal("Wanted", expected.Scope, "but got", actual.Scope)
	}

	if actual.ExpiresIn != expected.ExpiresIn {
		t.Fatal("Wanted", expected.ExpiresIn, "but got", actual.ExpiresIn)
	}

	if actual.TokenType != expected.TokenType {
		t.Fatal("Wanted", expected.TokenType, "but got", actual.TokenType)
	}
}
