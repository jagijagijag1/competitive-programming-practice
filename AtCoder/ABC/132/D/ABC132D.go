package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

var c [4005][4005]int

func main() {
	c[0][0] = 1
	for i := 0; i <= 4000; i++ {
		for j := 0; j <= i; j++ {
			c[i+1][j] += c[i][j] % mod
			c[i+1][j+1] += c[i][j] % mod
		}
	}

	sc.Split(bufio.ScanWords)
	N, K := nextInt(), nextInt()
	for i := 1; i <= K; i++ {
		b := f(K, i)
		b %= mod

		r := f2(N-K-(i-1), i+1)
		r %= mod
		// r := f(N-K, i-1)
		// r %= mod
		// r += f(N-K, i)
		// r %= mod
		// r += f(N-K, i)
		// r %= mod
		// r += f(N-K, i+1)
		// r %= mod

		res := (b * r) % mod
		fmt.Println(res)
	}
}

func f(n, k int) int {
	if n < k {
		return 0
	}
	if n == 0 && k == 0 {
		return 1
	}
	if k < 1 {
		return 0
	}
	return f2(n-k, k)
}

func f2(n, k int) int {
	return comb(n+k-1, k-1)
}

func comb(n, k int) int {
	return c[n][k]
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
