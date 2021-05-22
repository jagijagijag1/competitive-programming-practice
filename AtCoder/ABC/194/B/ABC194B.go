package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, B, mina, minb := nextInt(), []int{}, []int{}, 0, 0
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
		B = append(B, nextInt())

		if A[i] < A[mina] {
			mina = i
		}
		if B[i] < B[minb] {
			minb = i
		}
	}

	if mina != minb {
		fmt.Println(max(A[mina], B[minb]))
	} else {
		sort.Ints(A)
		sort.Ints(B)
		fmt.Println(min(A[0]+B[0], min(max(A[0], B[1]), max(A[1], B[0]))))
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
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

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
