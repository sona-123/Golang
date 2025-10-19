package main

import "fmt"

func vals(a int, b int) (int, int) {
	return a + b, a - b
}
func main() {
	a, b := vals(10, 40)
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals(20, 40)
	fmt.Println(c)
}
