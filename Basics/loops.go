package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			fmt.Println(j)
		}
	}

	//Range
	//Can't use range independently => list:=range 2
	nums := []int{1, 3, 2, 4, 8}
	for index, value := range nums {
		fmt.Println(index)
		fmt.Println(value)
	}

	for _, value := range nums {
		if value == 3 {
			break
		}
		fmt.Println(value)
	}

	for _, value := range "sonali" {
		if value == 'l' {
			break
		}
		fmt.Println(value)
	}

	strings := []string{"sonali", "key"}
	for _, value := range strings {
		fmt.Println(value)
	}
}
