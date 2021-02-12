package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K, hands := nextInt(), nextInt(), map[rune]int{}
	hands['s'], hands['p'], hands['r'] = nextInt(), nextInt(), nextInt()
	T := []rune(nextLine())

	h, res := make([]string, N), 0
	for i := 0; i < N; i++ {
		if i < K {
			res += hands[T[i]]
			h[i] = "win"
		} else {
			if T[i-K] == T[i] && h[i-K] == "win" {
				h[i] = "draw"
			} else {
				res += hands[T[i]]
				h[i] = "win"
			}
		}
	}

	fmt.Println(res)
}

type tuple struct {
	hand rune
	val  int
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
