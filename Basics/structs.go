package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main() {
	fmt.Println(person{"Sonali", 20})
	fmt.Println(person{name: "Alice", age: 10})
	fmt.Println(person{name: "sonali"})
	fmt.Println(&person{name: ""})
	fmt.Println(newPerson("sonali"))

	s := person{name: "Aditya", age: 60}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Tommy",
		true,
	}
	fmt.Println(dog)

}
