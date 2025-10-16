package main

import (
	"fmt"
	"math"
)

const s string = "Sonali"

func main() {
	fmt.Println(s)
	const num = 700
	const deno = 3e20 / num
	fmt.Println(deno)
	fmt.Println(deno)
	fmt.Println(math.Sincos(num))
}
