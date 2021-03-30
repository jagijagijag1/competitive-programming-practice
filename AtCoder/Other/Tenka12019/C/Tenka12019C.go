package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, S := nextInt(), []rune(nextLine())

	csb, csw := make([]int, N+1), make([]int, N+1)
	for i := 0; i < N; i++ {
		if S[i] == '#' {
			csb[i+1] = csb[i] + 1
		} else {
			csb[i+1] = csb[i]
		}
	}
	for i := N - 1; 0 <= i; i-- {
		if S[i] == '.' {
			csw[i] = csw[i+1] + 1
		} else {
			csw[i] = csw[i+1]
		}
	}

	res := 1000000000
	for i := 0; i < N+1; i++ {
		res = min(res, csb[i]+csw[i])
	}
	fmt.Println(res)
}

func ch(cond bool, ok, ng int) int {
	if cond {
		return ok
	}
	return ng
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// var sc = bufio.NewScanner((os.Stdin))

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
