package main

import "fmt"

//main function will get automatically called when you start the application
func main() {
	//var variable_name datatype
	var data int
	data = 1000
	//way 2 : Declaration of variable without declaring any type (Auto declaration)
	//variable:=value
	datal := 10

	//We can't do data:-100 .bcz we are just reassigning the variable
	fmt.Println(data)
	fmt.Println(datal)

	//In general, the datatypes we are having in golang are :-
	//Golang is case-sensitive
	//It is a functional coding language
	//1.) Primitive
	// int (all variants), float, string, bool, complex, uint, byte
	//2.) Non-Primitive
	// struct, map, chan, interface, array, slice (dynamic size array), rune, reflect, pointer

	name := `sonali`
	fmt.Println(name)
}
