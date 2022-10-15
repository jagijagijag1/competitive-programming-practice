package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	k, mod := nextInt(), 1000000007

	if k%9 != 0 {
		fmt.Println(0)
		return
	}

	dp := make([]int, k+1)
	dp[0] = 1
	for i := 1; i <= k; i++ {
		for j := 1; j <= 9; j++ {
			if 0 <= i-j {
				dp[i] += dp[i-j] % mod
			}
			dp[i] %= mod
		}
	}
	fmt.Println(dp[k])
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
