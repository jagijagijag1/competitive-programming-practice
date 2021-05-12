package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, res, mod := nextInt(), 3, 1000000007
	x, y, z := make([]int, N), make([]int, N), make([]int, N)
	x[0], y[0], z[0] = 1, 0, 0
	if nextInt() != 0 {
		fmt.Println(0)
		return
	}
	for i := 1; i < N; i++ {
		A, cnt := nextInt(), 0
		if x[i-1] == A {
			cnt++
		}
		if y[i-1] == A {
			cnt++
		}
		if z[i-1] == A {
			cnt++
		}
		res *= cnt
		res %= mod

		if x[i-1] == A {
			x[i] = x[i-1] + 1
			y[i] = y[i-1]
			z[i] = z[i-1]
		} else if y[i-1] == A {
			x[i] = x[i-1]
			y[i] = y[i-1] + 1
			z[i] = z[i-1]
		} else {
			x[i] = x[i-1]
			y[i] = y[i-1]
			z[i] = z[i-1] + 1
		}
	}
	fmt.Println(res)
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
