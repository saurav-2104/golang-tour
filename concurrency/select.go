package main

import (
	"time"
	"fmt"
)

func main() {
	 ch1 := make(chan string)
	 ch2 := make(chan string)
	 go func() {
	 	time.Sleep(2 * time.Second)
	 	ch1 <- "Saurav"
	 }()
	 go func() {
	 	ch2 <- "Gupta"
	 }()
	 for i := 0; i<2; i++{
	 	select {
	 		case msg1 := <- ch1:
	 			fmt.Printf("Message received from channel 1: %v\n", msg1)
	 		case msg2 := <- ch2:
	 			fmt.Printf("Message received from channel 2: %v\n", msg2)
	 	}
	 }
}