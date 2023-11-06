package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/petenilson/bravo"
	"github.com/petenilson/bravo/oauth"
	"github.com/petenilson/bravo/strava"
)

type StravaClient struct {
	bravo.HttpClient
	base          string
	authenticated bool
	clientSecret  string
	clientId      string
	Athlete       *bravo.Athlete
	AuthService   bravo.AuthService
}

func NewClient(athlete *bravo.Athlete, auth_service bravo.AuthService) *StravaClient {
	return &StravaClient{
		Athlete:     athlete,
		AuthService: auth_service,
		base:        "http://www.strava.com/api/v3",
	}
}

func (c *StravaClient) Authenticate(ctx context.Context, req *http.Request) error {
	// Get Auth object for Athlete
	var access_token string
	// Is Access Token Valid?
	if a := c.Athlete.Auth; a.ExpiresAt.Before(time.Now()) {
		// If no then request a new access token
		t, err := c.RefreshToken(
			ctx,
			a.RefreshToken,
			c.clientSecret,
			c.clientId,
			"refresh_token",
		)
		if err != nil {
			return err
		}
		access_token = t.AccessToken
		// Update our token
		if err := c.AuthService.Update(
			a.Id,
			&bravo.AuthUpdate{
				AccessToken:  t.AccessToken,
				RefreshToken: t.RefreshToken,
				ExpiresAt:    time.Unix(t.ExpiresAt, 0),
			},
		); err != nil {
			return err
		}
	} else {
		access_token = a.AccessToken
	}
	// Add access token to header
	req.Header.Add("Authorization", "Bearer "+access_token)
	return nil
}

func (c *StravaClient) GetEfforts(ctx context.Context, segment_id int, start, end time.Time) ([]*strava.Effort, error) {
	path := "/segment_efforts"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+path, nil)
	if err != nil {
		return nil, err
	}
	if c.Authenticate(ctx, req); err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("GetEfforts Resp: %d", resp.StatusCode))
	}
	effort := []*strava.Effort{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &effort); err != nil {
		return nil, err
	}
	return effort, nil
}

func (c *StravaClient) RefreshToken(
	ctx context.Context, refresh_token, client_secret, client_id, grant_type string,
) (*oauth.Token, error) {
	path := fmt.Sprintf(
		"/oauth/token?refresh_token=%s&client_secret=%s&client_id=%s&grant_type=%s",
		refresh_token,
		client_secret,
		client_id,
		grant_type,
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.base+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil,
			errors.New(
				fmt.Sprintf(
					"RefreshAccessToken: StravaAuthError: got %d, want: %d",
					resp.StatusCode,
					http.StatusOK,
				),
			)
	}
	token := oauth.Token{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &token); err != nil {
		return nil, err
	}
	return &token, nil
}
