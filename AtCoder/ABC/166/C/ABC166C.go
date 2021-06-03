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
	h, res := make([]int, n), make([]bool, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
		res[i] = true
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		if h[a] <= h[b] {
			res[a] = false
		}
		if h[b] <= h[a] {
			res[b] = false
		}
	}

	c := 0
	for i := 0; i < n; i++ {
		if res[i] {
			c++
		}
	}
	fmt.Println(c)
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
