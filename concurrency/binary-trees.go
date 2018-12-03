package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func _Walk(t *tree.Tree, ch chan int){
	if t == nil{
		return
	}
	_Walk(t.Left, ch)
	ch <- t.Value
	_Walk(t.Right, ch)
}


func Walk(t *tree.Tree, ch chan int){
	_Walk(t, ch)
	close(ch)
}


// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for val := range ch1{
		if val != <-ch2{
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	
}
