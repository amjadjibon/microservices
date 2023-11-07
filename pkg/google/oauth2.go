package google

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuth2Client struct {
	config *oauth2.Config
}

func NewOAuth2Client(
	clientID string,
	clientSecret string,
	redirectURL string,
	scopes []string) *OAuth2Client {
	return &OAuth2Client{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       scopes,
			Endpoint:     google.Endpoint,
		},
	}
}

// User represents the Google user information.
type User struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func (o *OAuth2Client) FetchUserInfo(token *oauth2.Token) (*User, error) {
	// Create an HTTP client with the OAuth2 token
	client := o.config.Client(context.Background(), token)

	// Make a GET request to the Google API to fetch user information
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("google API request failed with status code: %d", resp.StatusCode)
	}

	// Parse the response JSON into a User struct
	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (o *OAuth2Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return o.config.Exchange(ctx, code)
}
