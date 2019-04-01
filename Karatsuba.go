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

	return reverse(ans.String())
}

func minus(x, y string) string {
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

func Karatsuba(x, y string) string {
	x = fmt.Sprintf("%0*s", len(y), x)
	y = fmt.Sprintf("%0*s", len(x), y)

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
	adbc := minus(minus(aux, ac), bd)

	acShifted := ac + fmt.Sprintf("%0*d", acShift, 0)
	adbcShifted := adbc + fmt.Sprintf("%0*d", adbcShift, 0)

	return add(acShifted, add(adbcShifted, bd))
}
