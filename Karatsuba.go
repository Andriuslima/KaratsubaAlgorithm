package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isSmaller(u, v string) bool{
	nu := len(u)
	nv := len(v)

	if(nu < nv) { return true }

	if nu > nv { return false }

	for i := 0; i < nu; i++{
		if u[i] < v[i]{
			return true
		} else if u[i] > v[i] {
			return false
		}
	}
	return false
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func mult(x, y string) string {
	u, errorU := strconv.ParseUint(x, 10, 64)
	v, errorV := strconv.ParseUint(y, 10, 64)

	if errorU != nil {
		panic(errorU)
	}
	if errorV != nil {
		panic(errorU)
	}

	return strconv.Itoa(int(u * v))
}

func add(x, y string) string {
	x = fmt.Sprintf("%0*s", len(y), x)
	y = fmt.Sprintf("%0*s", len(x), y)

	var ans bytes.Buffer
	carry := 0
	sum := 0
	for i := len(x) - 1; i >= 0; i-- {
		u, errX := strconv.ParseInt(string(x[i]), 10, 64)
		v, errY := strconv.ParseInt(string(y[i]), 10, 64)
		check(errX)
		check(errY)

		sum = int(u + v + int64(carry))
		if sum < 10 {
			ans.WriteString(strconv.Itoa(sum))
			carry = 0
		} else {
			ans.WriteString(strconv.Itoa(sum % 10))
			carry = 1
		}
	}

	ans.WriteString(strconv.Itoa(carry))

	result, err := strconv.ParseInt(reverse(ans.String()), 10, 64)
	check(err)

	return strconv.Itoa(int(result))
}

func minusCast(x, y string) string {
	u, errorU := strconv.ParseUint(x, 10, 64)
	v, errorV := strconv.ParseUint(y, 10, 64)

	if errorU != nil {
		panic(errorU)
	}
	if errorV != nil {
		panic(errorU)
	}

	return strconv.Itoa(int(u - v))
}

func minus (x, y string) string {
	signal := "+"
	if isSmaller(x, y){
		aux := x
		x = y
		y = aux
		signal = "-"
	}

	//fmt.Printf("(%s - %s)\n", x, y)

	var ans bytes.Buffer
	nx := len(x)
	ny := len(y)
	diff := nx - ny
	var carry uint8

	for i := ny - 1; i >= 0; i--{
		aux1, err1 := strconv.ParseInt(string(x[i + diff]), 10, 64)
		aux2, err2 := strconv.ParseInt(string(x[i]), 10, 64)
		check(err1); check(err2)

		sub := aux1 - aux2 - int64(carry)

		if sub < 0{
			sub = sub+10
			carry = 1
		} else {
			carry = 0
		}

		ans.WriteString(strconv.Itoa(int(sub)))
	}

	for j := nx - ny - 1; j >= 0; j-- {
		if x[j] == '0' && carry > 0{
			ans.WriteString("9")
			continue
		}

		aux1, err1 := strconv.ParseInt(string(x[j]), 10, 64)
		check(err1)
		sub := aux1 - int64(carry)
		if j > 0 || sub > 0{
			ans.WriteString(strconv.Itoa(int(sub)))
		}
		carry = 0
	}

	ans.WriteString(signal)
	result, err := strconv.ParseInt(reverse(ans.String()), 10, 64)
	check(err)
	return strconv.Itoa(int(result))
}

func Karatsuba(x, y string) string {
	x = fmt.Sprintf("%0*s", len(y), x)
	y = fmt.Sprintf("%0*s", len(x), y)

	//fmt.Printf("Karatsuba for %s * %s\n", x, y)

	if len(x) != len(y) {
		msg := fmt.Sprintf("Wrong number format:x = %s\ny = %s", x, y)
		panic(msg)
	}

	n := len(x)
	n2 := uint(n / 2)
	acShift := n
	adbcShift := n2
	if n%2 != 0 {
		acShift += 1
		adbcShift += 1
	}

	if n == 1 {
		return mult(x, y)
	}

	a := x[:n2]
	b := x[n2:]
	c := y[:n2]
	d := y[n2:]

	ac := Karatsuba(a, c)
	bd := Karatsuba(b, d)
	aux := Karatsuba(add(a, b), add(c, d))
	adbc := minusCast(minusCast(aux, ac), bd)

	acShifted := ac + fmt.Sprintf("%0*d", acShift, 0)
	adbcShifted := adbc + fmt.Sprintf("%0*d", adbcShift, 0)

	return add(acShifted, add(adbcShifted, bd))
}
