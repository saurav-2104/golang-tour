package main

import (
	"fmt"
	"time"
	"sync"
)

type ThreadSafeCounter struct{
	hmap map[string]int
	mux sync.Mutex
}

func (counter *ThreadSafeCounter) Increment(key string) {
	counter.mux.Lock()
	counter.hmap[key]++
	counter.mux.Unlock()
}

func (counter *ThreadSafeCounter) getValue(key string) int {
	counter.mux.Lock()
	defer counter.mux.Unlock()
	return counter.hmap[key]
}

func main() {
	counter := ThreadSafeCounter{hmap: make(map[string] int)}
	for i := 0; i< 100; i++ {
		go counter.Increment("random")
		time.Sleep(2 * time.Millisecond)
	}
	fmt.Println(counter.getValue("random"))
}