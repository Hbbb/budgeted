package banks

import (
	"net/http"

	"github.com/plaid/plaid-go/plaid"
)

// BankClient wraps the Plaid API Client and stores the necessary parameters to create
// instances of the plaid.Client
type BankClient struct {
	ClientID    string
	Secret      string
	PublicKey   string
	plaidClient *plaid.Client
}

// NewBankClient initializes a BankClient and attaches a plaid.Client instance to it
// with the correct options.
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
