package main

import (
	"fmt"
	"github.com/uryyyyyyy/goSample/calc"
	"github.com/golang/glog"
)

func main() {
	var str = "world"
	fmt.Printf("hello, %s\n", str)
	var a int = calc.Sum(2, 7)
	glog.Info("Prepare to repel boarders")
	fmt.Printf("sum: %d\n", a)
}
