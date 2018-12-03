package main

import "fmt"

// Polymorphism, I guess
type Real interface {
	Abs() float64
}

type num float64

// It is implicit
func (f num) Abs() float64{
	if(f < 0){
		return -float64(f)
	}
	return float64(f)
}

func main() {
	f := num(-3.4)
	var inf = f
	// works even if f was nil. Java would throw a NullPointerException
	fmt.Println(inf.Abs())

	// empty interface.
	var i interface{}
	fmt.Printf("(%v : %T)\n", i, i)
	i = 42
	fmt.Printf("(%v : %T)\n", i, i)
}