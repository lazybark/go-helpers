package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// GetClient creates http.Client out of oauth config & token string
func GetClient(config *oauth2.Config, token string) (*http.Client, error) {
	tok, err := MakeToken(token)
	if err != nil {
		return nil, fmt.Errorf("[GetClient] unable to make token: %w", err)
	}
	return config.Client(context.Background(), tok), nil
}

// MakeToken converts config string to oauth token
func MakeToken(token string) (*oauth2.Token, error) {
	tok := &oauth2.Token{}
	err := json.Unmarshal([]byte(token), tok)
	return tok, err
}
