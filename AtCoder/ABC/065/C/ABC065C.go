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

	if 1 < abs(N-M) {
		fmt.Println(0)
		return
	}

	np := factorial(N) % mod
	mc := factorial(M) % mod
	if N == M {
		mc = (mc * 2) % mod
	}
	fmt.Println((np * mc) % mod)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const mod = 1000000007

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
			v %= mod
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
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
