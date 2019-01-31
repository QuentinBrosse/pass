package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type PasswordPrompt struct {
	// Password prompt label
	Label string

	// Password validation function
	Validate func(string) error

	password  string
	confirmed bool
}

// Run prompt the password and his confirmation
// It also handle retries and interrupt.
func (p *PasswordPrompt) Run() (string, error) {
	for !p.confirmed || p.password == "" {
		err := p.promptPassword()
		if err != nil {
			switch err {
			case ErrEOF:
				p.confirmed = false
				p.password = ""
				break
			case ErrInterrupt:
				return "", ErrAbort
			default:
				panic("prompt failure: " + err.Error())
			}
		}
	}
	return p.password, nil
}

// PromptPassword
func (p *PasswordPrompt) promptPassword() error {
	prompt := promptui.Prompt{
		Label:    p.Label,
		Mask:     '*',
		Validate: p.Validate,
	}

	isConfirmation := !p.confirmed && p.password != ""
	if isConfirmation {
		prompt.Label = p.Label + " confirmation"
		prompt.Validate = validatePasswordConfirmation(p.password)
	}

	password, err := prompt.Run()
	if err != nil {
		switch err {
		case promptui.ErrEOF, promptui.ErrInterrupt:
			return err
		default:
			panic("prompt failure: " + err.Error())
		}
	}

	if isConfirmation {
		p.confirmed = true
	}
	p.password = password

	return nil
}

// validatePasswordConfirmation check that the
// password confirmation is equal to the given password.
func validatePasswordConfirmation(password string) func(string) error {
	return func(passwordConfirmation string) error {
		if passwordConfirmation != password {
			return fmt.Errorf("the two passwords are not equals")
		}
		return nil
	}
}
