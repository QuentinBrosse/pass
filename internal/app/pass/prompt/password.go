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

	// Handle password confirmation
	Confirmation bool

	password  string
	aborted   bool
	confirmed bool
}

// Run prompt the password.
func (p *PasswordPrompt) Run() (string, error) {
	p.run()

	if p.aborted {
		return "", ErrAborted
	}
	return p.password, nil
}

func (p *PasswordPrompt) run() {
	err := p.promptPassword()
	if err != nil {
		switch err {
		case promptui.ErrEOF, promptui.ErrInterrupt:
			p.aborted = true
			return
		default:
			panic("prompt failure: " + err.Error())
		}
	}

	if p.Confirmation && !p.confirmed {
		p.run()
	}
}

func (p *PasswordPrompt) promptPassword() error {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   fmt.Sprintf("%s {{ . | green | bold }}: ", promptui.IconGood),
		Invalid: fmt.Sprintf("%s {{ . | red | bold }}: ", promptui.IconBad),
		Success: fmt.Sprintf("%s {{ . }}: ", promptui.IconGood),
	}

	prompt := promptui.Prompt{
		Label:     p.Label,
		Mask:      '*',
		Validate:  p.Validate,
		Templates: templates,
	}

	isConfirmation := !p.confirmed && p.password != ""
	if isConfirmation {
		prompt.Label = "Confirmation"
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
