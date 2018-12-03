package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y,2))
}


// To modify struct vaiables pointer receivers are required.
func (v *Vertex) Scale(scale float64){
	v.X *= scale
	v.Y *= scale
}
//In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
func ScaleFunc(v *Vertex, f float64) {
	v.X *=f
	v.Y *= f
}

func main() {
	v := Vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(v)
	ScaleFunc(&v, 20)
	fmt.Println(v)
}