package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var wrapper = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7D56F4")).
	Width(22).
	Align(lipgloss.Center)

var oddWordStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#04B575")).
	Padding(0, 1)

var evenWordStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#3C3C3C")).
	Background(lipgloss.Color("#FAFAFA")).
	Padding(0, 1)

const ver = `
   /\  
  /  \ `
const verClose = `
   /\  
  /__\ `
const rev = `
  \  / 
   \/  `
const revOpen = `
  \¯¯/ 
   \/  `

const cel = `
 /    \  
/      \ `
const celOpen = `
 /¯¯¯¯\  
/      \ `
const celClose = `
 /    \  
/______\ `

const lec = `
\      /
 \    / `
const lecClose = `
\      /
 \____/ `
const lecOpen = `
\¯¯¯¯¯¯/
 \    / `

func getWords(prompt string) []string {
	// Find all the possible words in the prompt
	regex := regexp.MustCompile(`ver|cel|rev|lec`)
	sliceResults := regex.FindAll([]byte(prompt), -1)

	// Convert the slice of bytes to a slice of strings
	wordList := make([]string, len(sliceResults))
	for i, word := range sliceResults {
		wordList[i] = string(word)
	}

	return wordList
}

func validatePrompt(prompt string) error {
	if strings.Trim(prompt, " ") == "" {
		return errors.New("please enter a valid prompt")
	}
	return nil
}

func drawPrompt(words []string) string {
	result := ``
	lastIndex := len(words) - 1

	for index, value := range words {
		switch string(value) {
		case "ver":
			if index == lastIndex {
				result += verClose
			} else {
				result += ver
			}

		case "cel":
			if index == lastIndex {
				result += celClose
			} else if index == 0 {
				result += celOpen
			} else {
				result += cel
			}

		case "rev":
			if index == 0 {
				result += revOpen
			} else {
				result += rev
			}

		case "lec":
			if index == lastIndex {
				result += lecClose
			} else if index == 0 {
				result += lecOpen
			} else {
				result += lec
			}
		}
	}

	return result
}

// drawListOfWords takes a list of words and draws them in a grid of 3 columns
func drawListOfWords(words []string) string {
	result := ``
	for index, word := range words {
		if index%3 == 0 && index != 0 {
			result += "\n"
		}

		if index%2 == 0 {
			result += oddWordStyle.Render(string(word))
		} else {
			result += evenWordStyle.Render(string(word))
		}

	}

	return result
}

var prompt string

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
				Value(&prompt).
				Validate(validatePrompt),
		))

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	words := getWords(prompt)

	if len(words) == 0 {
		log.Fatal("please enter a valid prompt")
	}

	fmt.Println(wrapper.Render(drawListOfWords(words)))
	fmt.Println(drawPrompt(words))

}
