package main

import (
	"fmt"
	"github.com/uryyyyyyy/goSample/calc"
)

func main() {
	var str = "world"
	fmt.Printf("hello, %s\n", str)
	var a int = calc.Sum(2, 7)
	fmt.Printf("sum: %d\n", a)
}
