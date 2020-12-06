package main

import "fmt"

func main() {
	alp := [4]string{"A", "B", "C", "D"}
	fmt.Println(alp)

	x := alp[0:2]
	fmt.Println(x)

	y := x[2:4]
	fmt.Println(y)

	z := y[0:1]
	fmt.Println(z)

	z[0] = "X"
	fmt.Println(alp)

	names := []string{}
	names = append(names, "Pla")
	fmt.Println(names)
	names = append(names, "Zen", "Ex")
	fmt.Println(names)
	human := []string{"Gon", "Ben", "Len"}
	names = append(names, human...)
	fmt.Println(names)
}
