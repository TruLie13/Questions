package main

import (
	"bufio"
	"example/hello/greetings" // Ensure this matches your module path
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Define the questions and corresponding handlers
	questions := []struct {
		prompt   string                      // Prompt for the question
		handler  func(string)                // Handler function for processing the answer
		validate func(string) (bool, string) // Validation function for the answer (optional)
	}{
		{
			prompt:  "Enter your name: ",
			handler: handleName,
		},
		{
			prompt:  "How old are you? ",
			handler: handleAge,
			validate: func(answer string) (bool, string) {
				age, err := strconv.Atoi(answer)
				ifNotPositiveIntError := "Please enter a number greater than 0."
				if err != nil || age <= 0 {
					return false, ifNotPositiveIntError
				}
				return true, ""
			},
		},
		// future questions
	}

	// Ask each question and handle the response
	for _, q := range questions {
		answer := askQuestion(scanner, q.prompt, q.validate)
		q.handler(answer)
	}
}

func askQuestion(scanner *bufio.Scanner, prompt string, validate func(string) (bool, string)) string {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		answer := scanner.Text()

		if validate != nil {
			valid, errMsg := validate(answer)
			if !valid {
				fmt.Println("Invalid input:", errMsg)
				continue
			}
		}

		return answer
	}
}

func handleName(name string) {
	fmt.Println(greetings.Hello(name))
}

func handleAge(ageStr string) {
	age, _ := strconv.Atoi(ageStr) // Safe to ignore error as input is validated
	fmt.Println(greetings.Age(age))
}
