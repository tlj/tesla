package tesla

import (
	"encoding/json"
	"time"
)

var (
	AuthURL      = "https://owner-api.teslamotors.com/oauth/token"
	clientID     = "e4a9949fcfa04068f59abb5a658f2bac0a3428e4652315490b659d5ab3f35a9e"
	clientSecret = "c75f14bbadc8bee3a7594412c31416f8300256d7668ea7e6e7f06727bfb9d220"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    int64  `json:"created_at"`
	Email        string `json:"email"`
}

func (t Token) ExpiresAt() int64 {
	return time.Unix(t.CreatedAt, 0).
		Add(time.Duration(t.ExpiresIn) * time.Second).
		Unix()
}

func (t Token) IsExpiring() bool {
	return time.Unix(t.ExpiresAt(), 0).Before(time.Now().Add(-2 * time.Hour))
}

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Client) Authorize(email, password string) (*Token, error) {
	auth := AuthRequest{
		GrantType:    "password",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Email:        email,
		Password:     password,
	}
	data, _ := json.Marshal(auth)
	body, err := c.post(AuthURL, data)
	if err != nil {
		return nil, err
	}
	t := &Token{}
	err = json.Unmarshal(body, t)
	if err != nil {
		return nil, err
	}
	t.Email = email
	c.Token = t
	return t, nil
}

func (c *Client) Reauthorize() (*Token, error) {
	auth := AuthRequest{
		GrantType:    "refresh_token",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: c.Token.RefreshToken,
	}
	data, _ := json.Marshal(auth)
	body, err := c.post(AuthURL, data)
	if err != nil {
		return nil, err
	}
	t := &Token{}
	err = json.Unmarshal(body, t)
	if err != nil {
		return nil, err
	}
	t.Email = c.Token.Email
	c.Token = t
	return t, nil
}

func (c *Client) RefreshToken(tokenChan chan Token) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			if c.Token.IsExpiring() {
				if _, err := c.Reauthorize(); err != nil {
					continue
				}
				tokenChan <- *c.Token

				ticker.Stop()
				ticker = time.NewTicker(time.Hour)
			}
		}
	}
}
