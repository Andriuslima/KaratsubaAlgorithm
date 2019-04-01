package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(Karatsuba(os.Args[1], os.Args[2]))

	// # Testing sum
	//x := 72274
	//y := 54732
	//fmt.Printf("%d + %d = %s\n", x, y, add(strconv.Itoa(x), strconv.Itoa(y)))
	//fmt.Printf("Expected: %d", x+y)

	x := os.Args[1]
	y := os.Args[2]
	fmt.Printf("%s * %s = %s", x, y, Karatsuba(x, y))
}
