package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	p, cnt := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}

	for i := 0; i < n; i++ {
		j := (p[i] - 1 - i + n) % n
		cnt[j%n]++
		cnt[(j+1)%n]++
		cnt[(j+2)%n]++
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, cnt[i])
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
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
