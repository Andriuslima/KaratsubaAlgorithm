package main

import (
	"fmt"
	"strconv"
)

func mult(x, y string) string{
	u, errorU := strconv.ParseUint(x, 10, 64)
	v, errorV := strconv.ParseUint(y, 10, 64)

	if errorU != nil{
		panic(errorU)
	}
	if errorV != nil{
		panic(errorU)
	}

	return strconv.Itoa(int(u * v))
}

func add(x, y string) string{
	u, errorU:= strconv.ParseUint(x, 10, 64)
	v, errorV := strconv.ParseUint(y, 10, 64)

	if errorU != nil{
		panic(errorU)
	}
	if errorV != nil{
		panic(errorU)
	}

	return strconv.Itoa(int(u + v))
}

func minus(x, y string) string {
	u, errorU := strconv.ParseUint(x, 10, 64)
	v, errorV := strconv.ParseUint(y, 10, 64)

	if errorU != nil{
		panic(errorU)
	}
	if errorV != nil{
		panic(errorU)
	}

	return strconv.Itoa(int(u - v))
}

func Karatsuba(x, y string) string {
	x = fmt.Sprintf("%0*s", len(y), x)
	y = fmt.Sprintf("%0*s", len(x), y)

	if len(x) != len(y){
		msg := fmt.Sprintf("Wrong number format:x = %s\ny = %s", x, y)
		panic(msg)
	}

	n := len(x)
	n2 := uint(n/2)
	acShift := n
	adbcShift := n2
	if n % 2 != 0 {
		acShift += 1
		adbcShift += 1
	}

	if n == 1{ return mult(x, y) }

	a := x[:n2]
	b := x[n2:]
	c := y[:n2]
	d := y[n2:]

	ac := Karatsuba(a, c)
	bd := Karatsuba(b, d)
	aux := Karatsuba(add(a, b), add(c, d))
	adbc := minus( minus(aux ,ac), bd)

	acShifted := ac+fmt.Sprintf("%0*d", acShift, 0)
	adbcShifted := adbc +fmt.Sprintf("%0*d", adbcShift, 0)

	return add(acShifted, add(adbcShifted, bd))
}