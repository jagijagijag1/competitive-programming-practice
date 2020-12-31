package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()

	A := make([][]int, N)
	for i := range A {
		K := nextInt()
		A[i] = make([]int, K)
		for j := range A[i] {
			A[i][j] = nextInt()
		}
	}

	fmt.Printf("%d\n", abc118B(M, A))
}

func abc118B(M int, A [][]int) (res int) {
	likes := make([]int, M)
	for _, a := range A {
		for _, aa := range a {
			likes[aa-1]++
		}
	}

	for _, l := range likes {
		if l == len(A) {
			res++
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
