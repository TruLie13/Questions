package greetings

import "fmt"
import "strings"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := "I was hoping you were JZ!"
	if strings.ToUpper(name) == "JZ" {
		message = " Welcome!"
	}
	return fmt.Sprintf("hi %v, ", name) + message
}

func Age(age int) string {
	return fmt.Sprintf("You are %v years old, ", age)
}
