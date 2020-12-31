package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC139B() {
	sc.Split(bufio.ScanWords)
	A, B := nextInt(), nextInt()

	fmt.Printf("%d\n", ABC139B(A, B))
}

// ABC139B ...
func ABC139B(A, B int) (res int) {
	if B == 1 {
		return 0
	}

	num := A
	res++
	for num < B {
		num += A - 1
		res++
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
