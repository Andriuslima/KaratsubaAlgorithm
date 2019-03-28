package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestKaratsuba(t *testing.T){
	fptr := flag.String("fpath", "resources/t01.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	check(err)

	s := bufio.NewScanner(f)
	for s.Scan(){
		args := strings.Fields(s.Text())
		x, errX := strconv.ParseInt(args[0], 10, 64)
		y, errY := strconv.ParseInt(args[1], 10, 64)
		check(errX); check(errY)

		res := x * y
		kara := Karatsuba(args[0], args[1])
		if res != kara{
			t.Errorf("Karatsuba was incorrect, got: %d, want: %d.", kara, res)
		}
	}
}
