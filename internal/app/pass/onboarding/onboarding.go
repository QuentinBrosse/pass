package onboarding

import (
	"fmt"
	"regexp"

	"github.com/QuentinBrosse/pass/internal/app/pass/keyring"
	"github.com/QuentinBrosse/pass/internal/app/pass/prompt"
	"github.com/fatih/color"
)

const passAsciiArt = ` ____   __   ____  ____ 
(  _ \ / _\ / ___)/ ___)
 ) __//    \\___ \\___ \
(__)  \_/\_/(____/(____/
`

const message = `Welcome to pass!

Before starting, please create a master password.
It will be used to encrypt all your other passwords, and must:
- be at least 8 characters
- contains at least one special character
- contains at least one lowercase and one uppercase letter
`

// Run the on boarding process.
func Run() error {
	password := keyring.GetMasterPassword()
	if password != "" {
		return nil
	}
	printOnBoardingMessage()

	password, err := promptMasterPassword(true)
	if err != nil {
		return err
	}

	keyring.SetMasterPassword(password)
	return nil
}

func printOnBoardingMessage() {
	fmt.Printf("%s%s", color.BlueString(passAsciiArt), message)
}

// promptMasterPassword prompts the master password and his confirmation.
func promptMasterPassword(creation bool) (string, error) {
	passwordPrompt := prompt.PasswordPrompt{
		Label: "Master password",
	}

	if creation {
		passwordPrompt.Validate = validateMasterPassword
		passwordPrompt.Confirmation = true
	}

	masterPassword, err := passwordPrompt.Run()
	if err != nil {
		return "", err
	}
	return masterPassword, nil
}

var specialCharRegex = regexp.MustCompile("[[:punct:]]+")
var lowercaseRegex = regexp.MustCompile("[[:lower:]]+")
var uppercaseRegex = regexp.MustCompile("[[:upper:]]+")

// validateMasterPassword validate the password complexity.
// Its length must be equal or greater than 8 and contain at least:
// one lowercase letter, one uppercase letter, one number
// and one special character.
func validateMasterPassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("master password must be at least 8 characters")
	}
	if !specialCharRegex.MatchString(password) {
		return fmt.Errorf("master password must contains at least one special character")
	}
	if !lowercaseRegex.MatchString(password) {
		return fmt.Errorf("master password must contains at least one lowercase letter")
	}
	if !uppercaseRegex.MatchString(password) {
		return fmt.Errorf("master password must contains at least one uppercase letter")
	}
	return nil
}
