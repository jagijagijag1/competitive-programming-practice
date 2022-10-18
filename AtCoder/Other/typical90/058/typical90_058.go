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
	mod := 100000

	results := make([]int, mod)
	for i := 0; i < mod; i++ {
		results[i] = (i + dsum(i)) % mod
	}

	ts := make([]int, mod)
	for i := 0; i < mod; i++ {
		ts[i] = -1
	}

	pos, cnt := n, 0
	for ; ts[pos] == -1; cnt++ {
		ts[pos] = cnt
		pos = results[pos]
	}

	cycle := cnt - ts[pos]
	if ts[pos] <= k {
		k = (k-ts[pos])%cycle + ts[pos]
	}

	for i := 0; i < mod; i++ {
		if ts[i] == k {
			fmt.Println(i)
		}
	}
}

func dsum(n int) (res int) {
	for 0 < n {
		res += n % 10
		n /= 10
	}
	return
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
