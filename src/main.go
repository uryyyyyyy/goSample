package main

import "fmt"

func main() {
	var str string = "world"
	fmt.Printf("hello, %s\n", str)
	var a int = sum(2, 7)
	fmt.Printf("sum: %d\n", a)
}

func sum(a int, b int) int {
	return a + b
}