package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println(Karatsuba(os.Args[1], os.Args[2]))
}
