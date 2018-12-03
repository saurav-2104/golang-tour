package main

import (
	"fmt"
)

func main() {
	hmap := make(map[string]Vertex)
	hmap["Jamshedpur"] = Vertex{127.001, 0.01}
	fmt.Println(hmap["Jamshedpur"])
	val, ok := hmap["Ranchi"]
	fmt.Println(val, ok)
}

type Vertex struct {
	lat, long float64 
}