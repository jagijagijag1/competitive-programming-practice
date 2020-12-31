package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC052B() {
	sc.Split(bufio.ScanWords)
	_, S := nextInt(), nextLine()

	fmt.Printf("%d\n", ABC052B(S))
}

// ABC052B ...
func ABC052B(S string) (res int) {
	curr := 0
	for _, s := range S {
		if s == 'I' {
			curr++
		} else {
			curr--
		}
		if res < curr {
			res = curr
		}
	}

	return
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	l := nextLine()
	i, e := strconv.Atoi(l)
	if e != nil {
		panic(e)
	}
	return i
}
