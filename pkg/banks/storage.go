package banks

import (
	"bdgt/pkg/core"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type banks struct {
	Banks []bank `json:"banks"`
}

type bank struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

// Add adds a bank to the config store
func Add(bankName string, accessToken string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	banksList := append(config.Banks, bank{Name: bankName, AccessToken: accessToken})
	config.Banks = banksList

	err = setConfig(config)
	if err != nil {
		return err
	}

	return nil
}

// Remove removes a bank from the config store
func Remove(bankName string) error {
	config, err := getConfig()
	if err != nil {
		return err
	}

	for i, bank := range config.Banks {
		if bank.Name == bankName {
			config.Banks = append(config.Banks[:i], config.Banks[i+1:]...)
		}
	}

	err = setConfig(config)
	if err != nil {
		return err
	}

	return nil
}

func getConfig() (*banks, error) {
	cPath, err := core.ConfigPath()
	if err != nil {
		return nil, err
	}

	f, err := ioutil.ReadFile(cPath + "/banks.json")
	if err != nil {
		return nil, err
	}

	config := &banks{}
	json.Unmarshal(f, &config)

	return config, nil
}

func setConfig(config *banks) error {
	configBytes, err := json.Marshal(config)
	if err != nil {
		return errors.New("banks: writing config to file")
	}

	cPath, err := core.ConfigPath()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(cPath+"/banks.json", configBytes, 644)
	if err != nil {
		return err
	}

	return nil
}
