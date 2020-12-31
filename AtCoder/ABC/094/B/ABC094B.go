package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC094B() {
	sc.Split(bufio.ScanWords)
	_, M, X, A := nextInt(), nextInt(), nextInt(), []int{}
	for i := 0; i < M; i++ {
		A = append(A, nextInt())
	}

	fmt.Printf("%d\n", ABC094B(X, A))
}

// ABC094B ...
func ABC094B(X int, A []int) int {
	idx := sort.Search(len(A), func(i int) bool {
		return A[i] >= X
	},
	)

	return minInt(idx, len(A)-idx)
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
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
