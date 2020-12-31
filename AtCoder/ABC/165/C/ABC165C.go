package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC165C() {
	sc.Split(bufio.ScanWords)
	N, M, Q, abcds := nextInt(), nextInt(), nextInt(), []abcd{}
	for i := 0; i < Q; i++ {
		abcds = append(abcds, abcd{nextInt(), nextInt(), nextInt(), nextInt()})
	}

	fmt.Printf("%d\n", ABC165C(N, M, abcds))
}

type abcd struct {
	a, b, c, d int
}

var resABC165C int

// ABC165C ...
func ABC165C(N, M int, abcds []abcd) int {
	dfsABC165C(N, M, abcds, []int{1})
	return resABC165C
}

func dfsABC165C(N, M int, abcds []abcd, A []int) {
	if len(A) == N+1 {
		score := 0
		for _, t := range abcds {
			if A[t.b]-A[t.a] == t.c {
				score += t.d
			}
		}
		resABC165C = maxInt(resABC165C, score)
		return
	}

	A = append(A, A[len(A)-1])
	for A[len(A)-1] <= M {
		dfsABC165C(N, M, abcds, A)
		A[len(A)-1]++
	}
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
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
