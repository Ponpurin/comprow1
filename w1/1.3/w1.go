package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	startA := time.Now()
	var welc string
	for i := 0; i < 10000; i++ {
		welc += "a"
	}
	fmt.Println("a", time.Since(startA))

	startB := time.Now()
	var welc2 strings.Builder
	for i := 0; i < 10000; i++ {
		welc2.WriteString("b")
	}
	fmt.Println("b", time.Since(startB))
}
