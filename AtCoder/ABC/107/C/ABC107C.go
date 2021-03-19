package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K, x := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		x = append(x, nextInt())
	}

	m := 1000000000
	for i := 0; i <= N-K; i++ {
		l2r := abs(x[i+K-1]) + abs(x[i+K-1]-x[i])
		r2l := abs(x[i]) + abs(x[i+K-1]-x[i])
		m = min(m, min(l2r, r2l))
	}

	fmt.Println(m)
}

func abs(a int) int {
	if a < 0 {
		return -a
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
