package main

import "fmt"

type rect struct {
	w, h int
}

func (r *rect) area() int {
	return r.w * r.h
}
func (r *rect) perimeter() int {
	return 2 * (r.w + r.h)
}
func main() {
	r := rect{w: 10, h: 20}
	fmt.Println("area: ", r.area())
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perimeter())
}
