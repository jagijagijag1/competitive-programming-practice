package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC092B() {
	sc.Split(bufio.ScanWords)
	N, D, X, A := nextInt(), nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	fmt.Printf("%d\n", ABC092B(D, X, A))
}

// ABC092B ...
func ABC092B(D, X int, A []int) (res int) {
	for _, a := range A {
		for i := 1; i <= D; i += a {
			res++
		}
	}

	return res + X
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
