package main

import (
	"fmt"
	"log"
	"os"

	"celvercel/pkg/helpers"
	"celvercel/pkg/prompt"
	"celvercel/pkg/styles"

	"github.com/charmbracelet/huh"
)

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

	fmt.Println(styles.Wrapper.Render(prompt.DrawListOfWords(words)))
	fmt.Println(prompt.DrawPrompt(words))
}
