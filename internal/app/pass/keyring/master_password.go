package keyring

import (
	keyring "github.com/zalando/go-keyring"
)

const (
	keyringService  = "Pass Master"
	keyringUsername = "pass"
)

var ErrNotFound = keyring.ErrNotFound

// TODO: Remove me
const TmpPassword = "🤫"

// Get the master password in the OS keyring system.
func GetMasterPassword() string {
	password, err := keyring.Get(keyringService, keyringUsername)
	if err != nil {
		switch err {
		case keyring.ErrNotFound:
			return ""
		default:
			panic("cannot get the master password: " + err.Error())
		}
	}
	return password
}

// Set the master password in the OS keyring system.
func SetMasterPassword(password string) {
	err := keyring.Set(keyringService, keyringUsername, password)
	if err != nil {
		panic("cannot set the master password: " + err.Error())
	}
}
