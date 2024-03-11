package prompt

import (
	"celvercel/pkg/helpers"
	"celvercel/pkg/styles"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
)

func CreateForm() (string, []string) {
	var userPrompt string

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

	words := GetWords(userPrompt)

	if len(words) == 0 {
		fmt.Println(styles.ErrorStyle.Render("You must enter a valid prompt :("))
		os.Exit(1)
	}

	// get the result and remove the first \n
	result := helpers.TrimFirstRune(DrawPrompt(words))

	fmt.Println(styles.Wrapper.Render(DrawListOfWords(words)))
	fmt.Println(result)

	return result, words
}
