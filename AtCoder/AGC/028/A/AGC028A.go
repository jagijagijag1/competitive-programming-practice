package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M, S, T := nextInt(), nextInt(), nextLine(), nextLine()

	xlen := lcm(N, M)
	ln, lm := xlen/N, xlen/M

	for nn, mm := 0, 0; nn < N && mm < M; {
		nidx, midx := nn*ln+1, mm*lm+1
		if nidx == midx {
			if S[nn] != T[mm] {
				fmt.Println(-1)
				return
			}
		}

		if nidx < midx {
			nn++
		} else {
			mm++
		}
	}

	fmt.Println(xlen)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
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
