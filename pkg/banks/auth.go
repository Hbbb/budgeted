package banks

import (
	"errors"
)

// ExchangePublicToken exchanges a public token for a secret access token
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
