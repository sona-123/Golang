package main

import "fmt"

func main() {
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = (i + j)
		}
	}
	fmt.Println(twoD)
}
