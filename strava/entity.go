package strava

import "time"

type Effort struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	ElapsedTime    int       `json:"elapsed_time"`
	StartDateLocal time.Time `json:"start_data_local"`
}

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}
