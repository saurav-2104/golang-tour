package main

import (
	"fmt"	
)

func add(arr [] int, c chan int) {
	sum := 0
	for _, v := range arr{
		sum += v
	}
	c <- sum
}

func main() {
	nums := []int{1,2,4,5,-3,-7,9}
	c := make(chan int)
	go add(nums[len(nums) / 2:], c)
	go add(nums[:len(nums) / 2], c)
	x, y := <-c, <-c
	fmt.Println(x+y)
}