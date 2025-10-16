package main

import (
	"fmt"
	"maps"
)

func main() {
	//map_name := make(map[key_name]value_name)
	mp := make(map[string]int)
	mp["sonali"] = 1
	mp["nandini"] = 2
	fmt.Println("map:", mp)

	value1 := mp["sonali"]
	fmt.Println(value1)

	//If value for a key doesn't exist, it will return 0
	fmt.Println(mp["key3"])

	//Length of the map
	fmt.Println("len:", len(mp))

	//Deleting a key
	delete(mp, "sonali")
	//Cleaning all entry
	clear(mp)

	//Check if a key exists
	_, prs := mp["nandini"]
	fmt.Println("prs:", prs)

	//Declaring a map literal
	n := map[string]int{"foo": 1, "bar": 2, "table": 3, "tennis": 4}
	fmt.Println(n)

	//Comparing maps
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n==n2")
	}
}
