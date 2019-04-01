package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestKaratsuba(t *testing.T){
	testFile := os.Args[2]
	fptr := flag.String("fpath", testFile, "file path to read from")
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
		kara, err := strconv.ParseInt(Karatsuba(args[0], args[1]), 10, 64)
		check(err)
		if res != kara{
			t.Errorf("Karatsuba was incorrect, got: %d, want: %d.", kara, res)
		}
	}
}
