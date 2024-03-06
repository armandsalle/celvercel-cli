package prompt

import (
	"celvercel/pkg/styles"
	"regexp"
)

func GetWords(prompt string) []string {
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

func DrawPrompt(words []string) string {
	result := ``
	lastIndex := len(words) - 1

	for index, value := range words {
		switch string(value) {
		case "ver":
			if index == lastIndex {
				result += VerClose
			} else {
				result += Ver
			}

		case "cel":
			if index == lastIndex {
				result += CelClose
			} else if index == 0 {
				result += CelOpen
			} else {
				result += Cel
			}

		case "rev":
			if index == 0 {
				result += RevOpen
			} else {
				result += Rev
			}

		case "lec":
			if index == lastIndex {
				result += LecClose
			} else if index == 0 {
				result += LecOpen
			} else {
				result += Lec
			}
		}
	}

	return result
}

// drawListOfWords takes a list of words and draws them in a grid of 3 columns
func DrawListOfWords(words []string) string {
	result := ``
	for index, word := range words {
		if index%3 == 0 && index != 0 {
			result += "\n"
		}

		if index%2 == 0 {
			result += styles.OddWordStyle.Render(string(word))
		} else {
			result += styles.EvenWordStyle.Render(string(word))
		}

	}

	return result
}
