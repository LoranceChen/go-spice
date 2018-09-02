package main

import (
	"github.com/lorancechen/go-spice/hello/subpkg"
	"fmt"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	a := subpkg.Name()
	fmt.Println(a)
	Sum(5, 5)
}
