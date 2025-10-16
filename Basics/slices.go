package main

import "fmt"

func main() {
	var s []string
	fmt.Println("unint:", s, s == nil, len(s) == 0)
	s = make([]string, 3)

	// 	arr := [3]int{1, 2, 3} // fixed length array
	// s := []int{1, 2, 3}    // slice (can grow or shrink)
	s[0] = "10"
	s[1] = "20"
	fmt.Println(s[1]) //20
	s = append(s, "30")
	s = append(s, "40", "50")
	fmt.Println(s)

	//Copying Slices
	copySlice := make([]string, len(s))
	copy(copySlice, s)
	fmt.Println(copySlice)

	//Slicing Parts
	nums := []int{10, 20, 30, 40, 50}
	part := nums[1:4]
	fmt.Println(part)

	fmt.Println(nums[:3]) //[10,20,30]
	fmt.Println(nums[2:]) //[30,40,50]
	fmt.Println(nums[:])  //whole slice
	//slices share memory
	nums1 := []int{1, 2, 3, 4}
	part1 := nums1[1:3]
	part1[0] = 99
	fmt.Println(nums1)
	fmt.Println(part1)

	//2D slice (slice of slices)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
