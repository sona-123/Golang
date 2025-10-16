package main

import "fmt"

/*
1. if-else statement
if(condition){
  //statements
}else{
  //statements
}

2. switch(data){
case var1:
	statement
case var2:
	statement
}
*/
func main() {
	// data := 10
	// fmt.Print(data)
	// fmt.Print("Enter a number")
	// var input int
	// fmt.Scanln(&input)
	// if input%2 == 0 {
	// 	fmt.Println("Hey You are even")
	// } else {
	// 	fmt.Println("Key is odd")
	// }

	// if x := 10; x%2 == 0 {
	// 	fmt.Println("Hey You are even")
	// } else {
	// 	fmt.Println("Key is odd")
	// }
	data := 100
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data * 90)
		fallthrough //Execute the next case also
	default:
		data = 9999
		fmt.Println(data)
	}
}
