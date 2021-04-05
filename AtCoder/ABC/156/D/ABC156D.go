package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, a, b := nextInt(), nextInt(), nextInt()
	mod := 1000000007

	res := modpow(2, n, mod) - 1
	res -= modchoose(n, a, mod)
	res -= modchoose(n, b, mod)
	for res < 0 {
		res += mod
	}
	fmt.Println(res)
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

func modchoose(n, a, p int) int {
	x, y := 1, 1

	for i := 0; i < a; i++ {
		x *= n - i
		x %= p
		y *= i + 1
		y %= p
	}

	inv := modpow(y, p-2, p)
	res := x * inv
	return res % p
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
