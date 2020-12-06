package main

import "fmt"

func main() {
	fmt.Printf("10 is of type %T\n", 10)
	fmt.Printf("10 is of type %T\n", string(10))
	fmt.Printf("10 is of type %T\n", float64(10))
}
