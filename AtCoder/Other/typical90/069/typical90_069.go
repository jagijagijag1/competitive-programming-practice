package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	mod := 1000000007

	if k == 1 {
		if n == 1 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else if n == 1 {
		fmt.Println(k % mod)
	} else if n == 2 {
		fmt.Println(k * (k - 1) % mod)
	} else {
		fmt.Println(k * (k - 1) % 1000000007 * modpow(k-2, n-2, 1000000007) % 1000000007)
	}
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
