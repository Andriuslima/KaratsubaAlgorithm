package main

import "fmt"

func main(){
	//fmt.Println(Karatsuba(os.Args[1], os.Args[2]))
	x := "72274"
	y := "54732"
	fmt.Printf("%s + %s = %s\n", x, y, string(add(x, y)))
}
