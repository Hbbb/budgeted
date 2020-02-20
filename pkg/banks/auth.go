package banks

import (
	"errors"
	"net/http"

	"github.com/plaid/plaid-go/plaid"
)

// BankClient exists
type BankClient struct {
	ClientID    string
	Secret      string
	PublicKey   string
	plaidClient *plaid.Client
}

// NewBankClient exists
func NewBankClient(clientID, secret, publicKey string) (*BankClient, error) {
	clientOptions := plaid.ClientOptions{
		ClientID:    clientID,
		Secret:      secret,
		PublicKey:   publicKey,
		Environment: plaid.Development,
		HTTPClient:  http.DefaultClient,
	}

	pc, err := plaid.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	bc := &BankClient{
		ClientID:    clientID,
		Secret:      secret,
		PublicKey:   publicKey,
		plaidClient: pc,
	}

	return bc, nil
}

// ExchangePublicToken exchanges a public token for a private token
func (bc *BankClient) ExchangePublicToken(publicToken string) (string, error) {
	if bc.plaidClient == nil {
		return "", errors.New("banks: must use NewBankClient initializer before using the BankClient")
	}

	resp, err := bc.plaidClient.ExchangePublicToken(publicToken)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
