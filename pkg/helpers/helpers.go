package helpers

import (
	"errors"
	"strings"
)

func ValidatePrompt(prompt string) error {
	if strings.Trim(prompt, " ") == "" {
		return errors.New("please enter a valid prompt")
	}
	return nil
}
