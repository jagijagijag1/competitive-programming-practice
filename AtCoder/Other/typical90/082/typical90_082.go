package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	mod := 1000000007
	l, r := nextInt(), nextInt()
	dl, dr := d(l), d(r)

	res := 0
	for i := dl; i <= dr; i++ {
		ll, rr := max(l, pow(10, i-1)), min(r, pow(10, i)-1)
		n := (rr - ll + 1) % mod
		s := (rr + ll) % mod
		res += div(n*s%mod, 2, mod) * i % mod
	}
	fmt.Println(res % mod)
}

func d(N int) int {
	return len(strconv.Itoa(N))
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func div(a, b, p int) int {
	return a * modpow(b, p-2, p) % p
}

func modpow(a, n, p int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a % p
	}
	if n%2 == 1 {
		return (a * modpow(a, n-1, p)) % p
	}
	t := modpow(a, n/2, p)
	return (t * t) % p
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
