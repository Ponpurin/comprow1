package main

import "fmt"

func main() {
	var a int
	a = 10
	fmt.Println(a)
	a = 20
	fmt.Println(a)
	const b = 10
	fmt.Println(b)
	//b = 20 ไม่สามารเปลี่ยนค่าแทน
	fmt.Println(b)
}
