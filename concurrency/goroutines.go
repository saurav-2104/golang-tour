package main

import (
	"fmt"
	"time"
)

func greet(str string) {
	for i:=0; i<5; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
	}
}

func main() {
	go greet("Saurav's thread")
	greet("Saurav")
}