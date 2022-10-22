package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
		y[i] = nextInt()
	}
	sort.Ints(x)
	sort.Ints(y)

	res := 0
	tmp1, tmp2 := 0, 0
	for i := 0; i < n; i++ {
		tmp1 += abs(x[i] - x[max(n/2-1, 0)])
		tmp2 += abs(x[i] - x[n/2])
	}
	res += min(tmp1, tmp2)

	tmp1, tmp2 = 0, 0
	for i := 0; i < n; i++ {
		tmp1 += abs(y[i] - y[max(n/2-1, 0)])
		tmp2 += abs(y[i] - y[n/2])
	}
	res += min(tmp1, tmp2)
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
