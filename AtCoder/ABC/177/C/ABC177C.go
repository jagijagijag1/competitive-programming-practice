package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, res, mod := nextInt(), 0, int(1e9+7)
	a, sum := make([]int, n), make([]int, n)
	a[0] = nextInt()
	sum[0] = a[0]
	for i := 1; i < n; i++ {
		a[i] = nextInt()
		sum[i] = sum[i-1] + a[i]
	}

	for i := 0; i < n; i++ {
		res += a[i] * ((sum[n-1] - sum[i]) % mod)
		res %= mod
	}
	fmt.Println(res)
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
