package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 10)
	}
	for j := 1; j < 10; j++ {
		dp[1][j] = 1
	}

	mod := 998244353
	for i := 2; i < n+1; i++ {
		for j := 1; j < 10; j++ {
			dp[i][j] = dp[i-1][j]
			if j == 1 {
				dp[i][j] += dp[i-1][j+1]
			} else if j == 9 {
				dp[i][j] += dp[i-1][j-1]
			} else {
				dp[i][j] += dp[i-1][j+1] + dp[i-1][j-1]
			}
			dp[i][j] %= mod
		}
	}

	res := 0
	for j := 1; j < 10; j++ {
		res += dp[n][j]
		res %= mod
	}
	fmt.Println(res)
}

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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

// var sc = bufio.NewScanner((os.Stdin))
