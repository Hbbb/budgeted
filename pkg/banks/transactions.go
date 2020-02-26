package banks

import (
	"errors"
	"fmt"

	"github.com/plaid/plaid-go/plaid"
)

// Transaction is the data we care about from Plaid
type Transaction struct {
	Name      string `json:"name"`
	Amount    string `json:"amount"`
	Date      string `json:"date"`
	AccountID string `json:"account_id"`
	City      string `json:"city"`
	ID        string `json:"id"`
}

// FetchTransactions fetches all transactions between two dates from the Plaid API. It will recursively
// fetch all transactions if there are more than 100 returned from the initial request.
func (bc *BankClient) FetchTransactions(startDate, endDate string) ([]Transaction, error) {
	if bc.plaidClient == nil {
		return nil, errors.New("banks: must use NewBankClient initializer before using the BankClient")
	}

	config, err := getConfig()
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for _, bank := range config.Banks {
		ts, err := bc.fetch(bank.AccessToken, startDate, endDate, transactions)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, ts...)
	}

	return transactions, nil
}

func (bc *BankClient) fetch(accessToken, startDate, endDate string, transactions []Transaction) ([]Transaction, error) {
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

	var newTrans []Transaction
	for _, tr := range resp.Transactions {
		t := Transaction{
			ID:        tr.ID,
			Name:      tr.Name,
			Amount:    fmt.Sprintf("%f", tr.Amount),
			Date:      tr.Date,
			AccountID: tr.AccountID,
			City:      tr.Location.City,
		}

		newTrans = append(newTrans, t)
	}

	transactions = append(transactions, newTrans...)

	if len(transactions) < resp.TotalTransactions {
		return bc.fetch(accessToken, startDate, endDate, transactions)
	}

	return transactions, nil
}
