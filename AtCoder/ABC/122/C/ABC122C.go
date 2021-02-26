package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, Q, S := nextInt(), nextInt(), nextLine()
	cs := make([]int, N)

	cs[0] = 0
	for i := 1; i < N; i++ {
		if S[i-1] == 'A' && S[i] == 'C' {
			cs[i] = cs[i-1] + 1
		} else {
			cs[i] = cs[i-1]
		}
	}

	for i := 0; i < Q; i++ {
		l, r := nextInt()-1, nextInt()-1
		fmt.Println(cs[r] - cs[l])
	}
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
