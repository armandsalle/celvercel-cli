package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/charmbracelet/huh"
)

var ver = `
                       ██                      
                      ████                     
                     ██████                    
                    ████████                   
                   ██████████                  
                  ████████████                 
                 ██████████████                
                ████████████████               
               ██████████████████              `

var rev = `
               ██████████████████              
                ████████████████               
                 ██████████████                
                  ████████████                 
                   ██████████                  
                    ████████                   
                     ██████                    
                      ████                     
                       ██                      `

var cel = `
              ████████████████████             
             ██████████████████████            
            ████████████████████████           
           ██████████████████████████          
          ████████████████████████████         
         ██████████████████████████████        
        ████████████████████████████████       
       ██████████████████████████████████      
      ████████████████████████████████████     `

var lec = `
      ████████████████████████████████████     
       ██████████████████████████████████      
        ████████████████████████████████       
         ██████████████████████████████        
          ████████████████████████████         
           ██████████████████████████          
            ████████████████████████           
             ██████████████████████            
              ████████████████████             `

func getWords(prompt string) ([][]byte, string) {
	// Find all the possible words in the prompt
	regex := regexp.MustCompile(`ver|cel|rev|lec`)
	sliceResults := regex.FindAll([]byte(prompt), -1)

	// Convert the byte slices to strings and join them
	wordList := make([]string, len(sliceResults))
	for i, word := range sliceResults {
		wordList[i] = string(word)
	}
	stringResult := strings.Join(wordList, " ")

	return sliceResults, stringResult
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
				Validate(func(prompt string) error {
					if strings.Trim(prompt, " ") == "" {
						return errors.New("please enter a valid prompt")
					}
					return nil
				})),
	)

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	words, wordString := getWords(prompt)

	if len(words) == 0 {
		log.Fatal("please enter a valid prompt")
	}

	result := ``

	for _, value := range words {
		switch string(value) {
		case "ver":
			result += ver
		case "cel":
			result += cel
		case "rev":
			result += rev
		case "lec":
			result += lec
		}
	}

	fmt.Println("Result for: ", wordString)
	fmt.Println(result)
}
