package banks

import (
	"errors"

	"github.com/plaid/plaid-go/plaid"
)

// FetchTransactions exists
func (bc *BankClient) FetchTransactions(startDate, endDate string) ([]plaid.Transaction, error) {
	if bc.plaidClient == nil {
		return nil, errors.New("banks: must use NewBankClient initializer before using the BankClient")
	}

	config, err := getConfig()
	if err != nil {
		return nil, err
	}

	var transactions []plaid.Transaction
	for _, bank := range config.Banks {
		ts, err := bc.fetch(bank.AccessToken, startDate, endDate, transactions)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, ts...)
	}

	return transactions, nil
}

func (bc *BankClient) fetch(accessToken, startDate, endDate string, transactions []plaid.Transaction) ([]plaid.Transaction, error) {
	opts := plaid.GetTransactionsOptions{
		StartDate:  startDate,
		EndDate:    endDate,
		AccountIDs: []string{},
		Count:      100,
		Offset:     len(transactions),
	}

	resp, err := bc.plaidClient.GetTransactionsWithOptions(accessToken, opts)
	if err != nil {
		return nil, err
	}

	transactions = append(transactions, resp.Transactions...)

	if len(transactions) < resp.TotalTransactions {
		return bc.fetch(accessToken, startDate, endDate, transactions)
	}

	return transactions, nil
}
