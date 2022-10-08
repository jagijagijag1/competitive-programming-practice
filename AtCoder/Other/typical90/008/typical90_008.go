package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, s := nextInt(), nextLine()
	dict := map[byte]int{}
	for i := range "atcoder" {
		dict["atcoder"[i]] = i
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 8)
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 8; j++ {
			dp[i+1][j] += dp[i][j]
			if idx, ok := dict[s[i]]; ok && j == idx {
				dp[i+1][j+1] += dp[i][j]
			}
		}
		for j := 0; j < 8; j++ {
			dp[i+1][j] %= 1000000007
		}
	}
	fmt.Println(dp[n][7])
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
	initialBufSize = 100000
	maxBufSize     = 10000000
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
