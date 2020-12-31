package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC007B() {
	sc.Split(bufio.ScanWords)
	A := nextLine()

	fmt.Printf("%s\n", ABC007B(A))
}

// ABC007B ...
func ABC007B(A string) string {
	if A == "a" {
		return "-1"
	}

	if len(A) == 1 {
		return string(byte(A[0]) - 1)
	}

	return A[0 : len(A)-1]
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
