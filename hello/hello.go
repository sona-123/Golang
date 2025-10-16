package main

import (
	"example.com/greetings/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)
	//message, err := greetings.Hello("Sonali")

	// Create a slice of names
	names := []string{"sonali", "rahul", "nandini", "tiya"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(message)
	fmt.Println(messages)
}
