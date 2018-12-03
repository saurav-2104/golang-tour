package main

import "fmt"

type Person struct {
	id int
	name string
}

func (p Person)String() string {
	ipAr = make([] int )
	return fmt.Sprintf("Name is %v and ID is %v", p.name, p.id)
}

func main() {
	saurav := Person{3, "Saurav"}
	fmt.Println(saurav)
}