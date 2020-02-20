package banks

import (
	"net/http"

	"github.com/plaid/plaid-go/plaid"
)

type AuthClient struct {
	ClientID, Secret, PublicKey string
}

func NewAuthClient(clientID, secret, publickey string) *AuthClient {
	return &AuthClient{ClientID: clientID, Secret: secret, PublicKey: publickey}
}

// ExchangePublicToken exchanges a public token for a private token
func (a *AuthClient) ExchangePublicToken(publicToken string) (string, error) {
	clientOptions := plaid.ClientOptions{
		ClientID:    a.ClientID,
		Secret:      a.Secret,
		PublicKey:   a.PublicKey,
		Environment: plaid.Development,
		HTTPClient:  http.DefaultClient,
	}

	client, err := plaid.NewClient(clientOptions)
	if err != nil {
		return "", err
	}

	resp, err := client.ExchangePublicToken(publicToken)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
