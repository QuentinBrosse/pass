package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

var confirmationLabel = "Confirmation"

// PasswordPrompt defines password Prompt configuration.
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

// Run runs the password prompt and return the password.
func (p *PasswordPrompt) Run() (string, error) {
	p.run()

	if p.aborted {
		return "", ErrAborted
	}
	return p.password, nil
}

// Run prompts the password and fill the PasswordPrompt struct.
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

// PromptPassword prompts a single password prompt for the password itself or for the confirmation.
func (p *PasswordPrompt) promptPassword() error {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   fmt.Sprintf("%s {{ . | green | bold }} : ", promptui.IconGood),
		Invalid: fmt.Sprintf("%s {{ . | red | bold }} : ", promptui.IconBad),
		Success: fmt.Sprintf("%s {{ . }} : ", promptui.IconGood),
	}

	prompt := promptui.Prompt{
		Label:     p.Label,
		Mask:      '*',
		Validate:  p.Validate,
		Templates: templates,
	}

	isConfirmation := !p.confirmed && p.password != ""
	if isConfirmation {
		prompt.Label = confirmationLabel
		prompt.Validate = validatePasswordConfirmation(p.password)
	}

	if p.Confirmation {
		prompt.Label = fmt.Sprintf("%-*s", p.getLabelPadding(), prompt.Label)
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

// GetLabelPadding returns the padding required to align labels.
func (p *PasswordPrompt) getLabelPadding() int {
	labelPadding := len(confirmationLabel)
	if len(p.Label) > labelPadding {
		return len(p.Label)
	}
	return labelPadding
}

// ValidatePasswordConfirmation checks that the
// password confirmation is equal to the given password.
func validatePasswordConfirmation(password string) func(string) error {
	return func(passwordConfirmation string) error {
		if passwordConfirmation != password {
			return fmt.Errorf("the two passwords are not equals")
		}
		return nil
	}
}