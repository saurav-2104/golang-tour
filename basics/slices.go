package main

import "fmt"

func main() {
	a := [5] int {0,1,2,3,4}
	fmt.Println(a)
	// Slice are reference to arrays
	s1 := a[:3]
	s2 := a[2:]
	fmt.Println(s1, s2)
	s1[2] = 546
	fmt.Println(a, s1, s2)
	// slice literal
	primes := []int{2,3,5,7,11}
	fmt.Println("Primes: ",primes)

	for i,v := range(primes){
		fmt.Println(i, v)
	}

	// len(s) is length of slice and cap(s) is the capacity of underlying array
	primes = primes[2:]
	printS(primes)
	primes = primes[:1]
	printS(primes)

	// creating slice dynamically using make
	slice := make([] int, 5)
	printS(slice)
}

func printS(x []int) {
	fmt.Println(x, len(x), cap(x))
}