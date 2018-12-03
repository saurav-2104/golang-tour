package main

import "fmt"

func jdi(i interface {}) {
	switch s := i.(type) {
	case int:
		fmt.Println("Looks like an integer")
	case string:
		fmt.Println("Looks like a string")
	default:
		fmt.Printf("No idea what the heck this is: %T\n", s)
	}
}

func main() {
	jdi(2)
	jdi("Saurav")
	jdi(false)
}