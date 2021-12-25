package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	k, a, b := nextInt(), nextInt(), nextInt()

	aa, bb, kk := 0, 0, 1
	for 0 < a || 0 < b {
		aa += a % 10 * kk
		bb += b % 10 * kk
		a, b, kk = a/10, b/10, kk*k
	}
	fmt.Println(aa * bb)
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

// var sc = bufio.NewScanner((os.Stdin))
