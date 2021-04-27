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
	dp := make([]int, N+1)
	dp[0] = 0

	for i := 1; i <= N; i++ {
		dp[i] = N
		for p := 1; p <= i; p *= 6 {
			dp[i] = min(dp[i], dp[i-p]+1)
		}
		for p := 1; p <= i; p *= 9 {
			dp[i] = min(dp[i], dp[i-p]+1)
		}
	}
	fmt.Println(dp[N])
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
