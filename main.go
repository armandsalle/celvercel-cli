package main

import (
	"fmt"
	"log"
	"os"
	"unicode/utf8"

	"celvercel/pkg/helpers"
	"celvercel/pkg/prompt"
	"celvercel/pkg/styles"

	"github.com/charmbracelet/huh"
)

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

var userPrompt string

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.
				NewInput().
				Title("What are we drawing?").
				Description("Possible values: ver | cel | rev | lec").
				Prompt("▲ ").
				CharLimit(100).
				Placeholder("lecrev").
				Value(&userPrompt).
				Validate(helpers.ValidatePrompt),
		))

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	words := prompt.GetWords(userPrompt)

	if len(words) == 0 {
		fmt.Println(styles.ErrorStyle.Render("You must enter a valid prompt :("))
		os.Exit(1)
	}

	// get the result and remove the first \n
	result := trimFirstRune(prompt.DrawPrompt(words))

	fmt.Println(styles.Wrapper.Render(prompt.DrawListOfWords(words)))
	fmt.Println(result)

	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(result)

	if err2 != nil {
		log.Fatal(err2)
	}

}
