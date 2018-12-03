package main

import "fmt"

type Vertex struct {
	id int
	name string
}

func main() {
	v := Vertex{1, "Saurav"}
	fmt.Println(Vertex{2, "Nilay"}, v)
}