package oauth

type OAuthService interface {
	RefreshToken(refresh_token, client_secret, client_id, grant_type string) (*Token, error)
}

type Token struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}
