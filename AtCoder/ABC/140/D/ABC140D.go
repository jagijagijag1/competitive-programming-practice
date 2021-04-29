package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K, S, res := nextInt(), nextInt(), nextLine(), 0
	for i := 0; i < N-1; i++ {
		if S[i] == S[i+1] {
			res++
		}
	}
	fmt.Println(min(res+2*K, N-1))
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
