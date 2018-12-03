package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello ! I am Saurav.")
	bytes := make([] byte, 8)
	for {
		n, err := r.Read(bytes)
		fmt.Printf("N is: %v and err is: %v\n", n, err)
		fmt.Printf("String formed: %q\n", bytes[:n])
		if err == io.EOF {
			break
		}
	}
}