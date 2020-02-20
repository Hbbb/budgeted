package banks

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type configStore struct {
	Banks []bankConfig `json:"banks"`
}

type bankConfig struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

// Add adds a bank to the config store
func Add(bankName string, accessToken string) error {
	f, err := ioutil.ReadFile("banks.json")
	if err != nil {
		return err
	}

	config := &configStore{}
	json.Unmarshal(f, &config)

	banksList := append(config.Banks, bankConfig{Name: bankName, AccessToken: accessToken})
	config.Banks = banksList

	configBytes, err := json.Marshal(config)
	if err != nil {
		return errors.New("banks: error adding new bank")
	}

	err = ioutil.WriteFile("banks.json", configBytes, 644)
	if err != nil {
		return err
	}

	return nil
}

// Remove removes a bank from the config store
func Remove(bankName string) error {
	return nil
}
