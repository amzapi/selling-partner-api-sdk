package oauth2

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	ScopeNotificationsApi = "sellingpartnerapi::notifications"
	ScopeMigrationApi     = "sellingpartnerapi::migration"
)

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
}

type Token struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
	ExpiresIn    int       `json:"expires_in"`
}

// expiryDelta determines how earlier a token should be considered
// expired than its actual expiration time. It is used to avoid late
// expirations due to client-server time mismatches.
const expiryDelta = 10 * time.Second

// timeNow is time.Now but pulled out as a variable for tests.
var timeNow = time.Now

// expired reports whether the token is expired.
// t must be non-nil.
func (t *Token) expired() bool {
	if t.Expiry.IsZero() {
		return false
	}
	return t.Expiry.Round(0).Add(-expiryDelta).Before(timeNow())
}

// Valid reports whether t is non-nil, has an AccessToken, and is not expired.
func (t *Token) Valid() bool {
	return t != nil && t.AccessToken != "" && !t.expired()
}

func RefreshAccessToken(cfg Config) (token *Token, err error) {
	reqBody, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	client := http.DefaultClient
	resp, err := client.Post("https://api.amazon.com/auth/o2/token", "application/json;charset=UTF-8", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respBody, &token)
	if err != nil {
		return nil, err
	}
	token.Expiry = time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	return token, nil
}
