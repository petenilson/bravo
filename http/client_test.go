package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/petenilson/bravo"
	"github.com/petenilson/bravo/mock"
)

func Do(req *http.Request) (*http.Response, error) {
	switch req.URL.Path {
	case "/api/v3/oauth/token":
		data, err := os.ReadFile("testdata/refresh_token.json")
		if err != nil {
			panic(err)
		}
		resp := http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(data)),
		}
		return &resp, nil
	case "/api/v3/segment_efforts":
		if auth := req.Header.Get("Authorization"); !strings.Contains(auth, "Bearer") {
			return &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": "auth details not provided"}`))),
			}, nil
		}
		data, err := os.ReadFile("testdata/segment_effort.json")
		if err != nil {
			panic(err)
		}
		resp := http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(data)),
		}
		return &resp, nil
	default:
		resp := http.Response{
			StatusCode: http.StatusNotFound,
			Body:       io.NopCloser(bytes.NewReader([]byte(`{"error": "not found"}`))),
		}
		return &resp, nil
	}
}

func NewTestClient(a *bravo.Athlete) *StravaClient {
	return &StravaClient{
		base:    "http://www.strava.com/api/v3",
		Athlete: a,
		HttpClient: &mock.HttpClient{
			DoFunc: Do,
		},
	}
}

func TestStravaClient(t *testing.T) {
	t.Run("GetEfforts with expired token", func(t *testing.T) {
		authService := mock.AuthService{}
		authService.UpdateFunc = func(i int, au *bravo.AuthUpdate) error {
			return nil
		}
		athlete := bravo.Athlete{
			Id:        1,
			StravaId:  12345,
			FirstName: "Peter",
			LastName:  "Nilson",
			Auth: &bravo.Auth{
				Id:           1,
				ExpiresAt:    time.Now().Add(-1 * time.Hour),
				AccessToken:  "dummy_access_token",
				RefreshToken: "dummy_refresh_token",
			},
		}
		client := NewTestClient(&athlete)
		client.AuthService = &authService
		_, err := client.GetEfforts(context.Background(), 1234, time.Now(), time.Now())
		if err != nil {
			t.Errorf("Got Error: %v", err)
		}
	})
	t.Run("GetEfforts with valid token", func(t *testing.T) {
		athlete := bravo.Athlete{
			Id:        1,
			StravaId:  12345,
			FirstName: "Peter",
			LastName:  "Nilson",
			Auth: &bravo.Auth{
				Id:           1,
				ExpiresAt:    time.Now().Add(1 * time.Hour),
				AccessToken:  "dummy_access_token",
				RefreshToken: "dummy_refresh_token",
			},
		}
		client := NewTestClient(&athlete)
		_, err := client.GetEfforts(context.Background(), 1234, time.Now(), time.Now())
		if err != nil {
			t.Fail()
		}
	})
}
