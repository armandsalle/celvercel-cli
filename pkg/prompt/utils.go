package prompt

import (
	"celvercel/pkg/styles"
	"regexp"
)

func GetWords(prompt string) []string {
	regex := regexp.MustCompile(`ver|cel|rev|lec`)
	sliceResults := regex.FindAll([]byte(prompt), -1)

	wordList := make([]string, len(sliceResults))
	for i, word := range sliceResults {
		wordList[i] = string(word)
	}

	return wordList
}

func getAsciiArt(index int, lastIndex int, nextWord string, asciiArt AsciiArt) string {
	// If there is only one word, return the only ascii art
	if index == lastIndex && index == 0 {
		return asciiArt.Only
	}

	// If it's the first word, return the open ascii art
	if index == 0 {
		openBeforeAscii := asciiArt.HandleOpenBefore(nextWord)

		if openBeforeAscii == "" {
			return asciiArt.Open
		}

		return openBeforeAscii
	}

	// If it's the last word, return the close ascii art
	if index == lastIndex {
		return asciiArt.Close
	}

	// If it's not the first or last word, return the default ascii art
	beforeAscii := asciiArt.HandleBefore(nextWord)
	if beforeAscii != "" {
		return beforeAscii
	}

	return asciiArt.Default
}

func DrawPrompt(words []string) string {
	result := ``
	lastIndex := len(words) - 1

	for index, value := range words {
		var nextWord string
		if index != lastIndex {
			nextWord = words[index+1]
		}

		switch string(value) {
		case "ver":
			result += getAsciiArt(index, lastIndex, nextWord, VerAscii)
		case "cel":
			result += getAsciiArt(index, lastIndex, nextWord, CelAscii)
		case "rev":
			result += getAsciiArt(index, lastIndex, nextWord, RevAscii)
		case "lec":
			result += getAsciiArt(index, lastIndex, nextWord, LecAscii)
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
