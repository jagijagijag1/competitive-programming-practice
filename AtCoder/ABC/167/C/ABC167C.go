package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m, x := nextInt(), nextInt(), nextInt()
	c := make([]int, n)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
	}

	res := 0
	compWithAll := make([]int, m)
	for i := 0; i < n; i++ {
		c[i] = nextInt()
		res += c[i]
		for j := 0; j < m; j++ {
			a[i][j] = nextInt()
			compWithAll[j] += a[i][j]
		}
	}
	for j := range compWithAll {
		if compWithAll[j] < x {
			fmt.Println(-1)
			return
		}
	}

	for bit := 0; bit < (1 << uint(n)); bit++ {
		cost, comp := 0, make([]int, m)
		for i := 0; i < n; i++ {
			if bit>>uint(i)&1 == 1 {
				cost += c[i]
				for j := 0; j < m; j++ {
					comp[j] += a[i][j]
				}
			}
		}

		allok := true
		for j := range comp {
			if comp[j] < x {
				allok = false
				break
			}
		}

		if allok {
			res = min(res, cost)
		}
	}
	fmt.Println(res)
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
