package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC088C() {
	sc.Split(bufio.ScanWords)
	X, Y := uint64(nextInt()), uint64(nextInt())

	fmt.Printf("%d\n", ARC088C(X, Y))
}

// ARC088C ...
func ARC088C(X, Y uint64) (res int) {
	for X <= Y {
		X = X * 2
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
