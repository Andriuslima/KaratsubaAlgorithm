package main

import (
	"fmt"
	"strconv"
)

func mult(x, y string) string{
	u, error_u := strconv.ParseUint(x, 10, 64)
	v, error_v := strconv.ParseUint(y, 10, 64)

	if error_u != nil{
		panic(error_u)
	}
	if error_v != nil{
		panic(error_u)
	}

	return string(u * v)
}

func add(x, y string) string{
	u, error_u := strconv.ParseUint(x, 10, 64)
	v, error_v := strconv.ParseUint(y, 10, 64)

	if error_u != nil{
		panic(error_u)
	}
	if error_v != nil{
		panic(error_u)
	}


	return string(u + v)
}

func minus(x, y string) string{
	u, error_u := strconv.ParseUint(x, 10, 64)
	v, error_v := strconv.ParseUint(y, 10, 64)

	if error_u != nil{
		panic(error_u)
	}
	if error_v != nil{
		panic(error_u)
	}

	return string(u - v)
}

func Karatsuba(x, y string) string {
	lenX := len(x)
	lenY := len(y)

	if lenX == 1 || lenY == 1{
		return mult(x, y)
	}

	a := x[:uint8(lenX/2)]
	b := x[uint8(lenX/2):]

	c := y[:uint8(lenY/2)]
	d := y[uint8(lenX/2):]

	fmt.Println(a, b, c, d)

	ac := Karatsuba(a, c)
	bd := Karatsuba(b, d)
	ad_bc := minus( minus( Karatsuba(add(a, b), add(c, d)) ,ac), bd)

	fmt.Println(a,b,c,d, ac, bd, ad_bc)

	return add(ac+fmt.Sprintf("%0*d", lenX-1, 0), add(ad_bc+fmt.Sprintf("%0*d", (lenX/2)-1, 0), bd))
}

func main(){
	fmt.Println(Karatsuba("30", "10"))
}