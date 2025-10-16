package main

import (
	"example.com/greetings/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Sonali")
	fmt.Println(message)
}
