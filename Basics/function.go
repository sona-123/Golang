package main

import (
	"fmt"
	"math"
)

/*func function_name(){}*/
func main() {
	result, x := solve()
	fmt.Println(result)
	fmt.Println(x)
	fmt.Println(math.Sqrt(100))
	// solve()
}
func solve() (int, string) {
	return 10, "sonali"
	// solve1()
	// fmt.Println(100)
}
func solve1(args ...int) {
	fmt.Println(args)
	solve2()
}
func solve2() (i int, j string) {
	i = 20
	j = "rahul"
	return
	//fmt.Println("Hello")
}
