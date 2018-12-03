package main

import "fmt"

func main() {
	add := ader()
	for i:=0; i<=10; i++ {
		fmt.Println(add(i))
	}
}

func ader() func(int) int  {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}