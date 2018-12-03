package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sep := strings.Fields(s)
	res := make(map[string]int)
	for _,val := range(sep){
			res[val]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
