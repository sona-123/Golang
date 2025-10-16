package main

import "fmt"

func main() {
	//Syntax : var array_name []type
	var arr [5]int
	fmt.Println("emp:", arr)

	arr[4] = 1000
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

}
