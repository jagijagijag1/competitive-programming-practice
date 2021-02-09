package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	A, ss := make([]int, N), make([]int, N)
	A[0] = nextInt()
	ss[0] = A[0]
	for i := 1; i < N; i++ {
		A[i] = nextInt()
		ss[i] = ss[i-1] + A[i]
	}

	res := abs(ss[0] - ss[N-1] + ss[0])
	for i := 1; i < N-2; i++ {
		x, y := ss[i], ss[N-1]-ss[i]
		res = min(res, abs(x-y))
	}

	fmt.Println(res)
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
