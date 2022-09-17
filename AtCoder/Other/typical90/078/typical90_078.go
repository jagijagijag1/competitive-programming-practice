package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	g := make([]int, n)

	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		if a < b {
			g[b]++
		}
		if b < a {
			g[a]++
		}
	}

	res := 0
	for i := range g {
		if g[i] == 1 {
			res++
		}
	}
	fmt.Println(res)
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
