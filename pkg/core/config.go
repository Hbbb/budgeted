package core

import "os/user"

// ConfigPath returns the path to the configuration folder
func ConfigPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir + "/.budgeted", nil
}
