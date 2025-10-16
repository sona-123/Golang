package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	//"golang.org/x/text/message"
)

// Syntax
// func Function_Name(Parameter_Name Parameter_Type) Return_Type

// Hello returns a greeting for a named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//message := fmt.Sprintf("Hi, %v. Welcome!", name)
	//For random string to get printed
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

/*
Print()	Prints values without formatting (no newline by default)
Println()	Prints values and adds a newline
Printf()	Prints formatted output (to console)
Sprintf()	Formats and returns a string
Fprintf()	Prints formatted output to an io.Writer (like a file or socket)
Errorf()	Creates a formatted error value
*/

// randomFormat returns one of a set of greeting messages. The returned message is selected at random
// Learn Slices
func randomFormat() string {
	formats := []string{
		"Hi %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

// Hellos returned a map that associates each of the named people
// Syntax :- func function_name(parameter_name parameter_type) (map[key_parameter]value_parameter, error)
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
