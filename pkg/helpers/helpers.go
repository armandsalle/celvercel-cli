package helpers

import (
	"errors"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func ValidatePrompt(prompt string) error {
	if strings.Trim(prompt, " ") == "" {
		return errors.New("please enter a valid prompt")
	}
	return nil
}

func TrimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func SaveResultToFile(fileName string, words []string, result string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer f.Close()

	_, err2 := f.WriteString(strings.Join(words, ", ") + "\n\n" + result)

	if err2 != nil {
		log.Fatal(err2)
		os.Exit(1)
	}
}
