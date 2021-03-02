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
	broken := make([]bool, N+1)
	for i := 0; i < M; i++ {
		broken[nextInt()] = true
	}

	fib, mod := make([]int, N+1), 1000000007
	fib[0] = 1
	for i := 0; i < N; i++ {
		for j := i + 1; j <= min(N, i+2); j++ {
			if !broken[j] {
				fib[j] += fib[i]
				fib[j] %= mod
			}
		}
	}
	fmt.Println(fib[N])
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
