package main

import (
	"fmt"
	"os"
)

func main() {
	x := os.Args[1]
	y := os.Args[2]
	fmt.Printf(Karatsuba(x, y))
}
