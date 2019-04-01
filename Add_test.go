package main

import (
	"bufio"
	"flag"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	testFile := os.Args[4]
	fptr := flag.String("fpath", testFile, "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	check(err)

	s := bufio.NewScanner(f)
	for s.Scan() {
		args := strings.Fields(s.Text())
		x, errX := strconv.ParseInt(args[0], 10, 64)
		y, errY := strconv.ParseInt(args[1], 10, 64)
		check(errX)
		check(errY)

		expected := x + y
		result, err := strconv.ParseInt(add(strconv.Itoa(int(x)), strconv.Itoa(int(y))), 10, 64)
		check(err)

		if expected != result {
			t.Errorf("Adition (%d + %d) was incorrect, got: %d, want: %d.", x, y, result, expected)
		} else {
			t.Logf("Adition (%d + %d) correct, got: %d", x, y, result)
		}
	}
}
