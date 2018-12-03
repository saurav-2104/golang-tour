package main

import "fmt"

func main() {
	var i interface {} = "saurav"
	val, ok := i.(string)
	fmt.Println(val, ok)
	in, ok := i.(int)
	fmt.Println(in, ok)

	// panic
	pan := i.(int)
	fmt.Println(pan)
}