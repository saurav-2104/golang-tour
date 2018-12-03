package main

import "fmt"

func main() {
	i := 2
	fmt.Println(i)
	//inc(i)
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}