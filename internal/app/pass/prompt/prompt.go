package prompt

import (
	"text/template"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func init() {
	// Redefine the FuncMap with github.com/fatih/color helpers
	promptui.FuncMap = template.FuncMap{
		"red":   color.RedString,
		"green": color.GreenString,
		"bold":  color.New(color.Bold).Sprintf,
	}
}
