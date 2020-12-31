package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC127C() {
	sc.Split(bufio.ScanWords)
	N, M, L, R := nextInt(), nextInt(), []int{}, []int{}
	for i := 0; i < M; i++ {
		L = append(L, nextInt())
		R = append(R, nextInt())
	}

	fmt.Printf("%d\n", ABC127C(N, L, R))
}

// ABC127C ...
func ABC127C(N int, L, R []int) int {
	lmax, rmin := 0, 10000000
	for i := range L {
		if lmax < L[i] {
			lmax = L[i]
		}
		if R[i] < rmin {
			rmin = R[i]
		}
	}

	if rmin < lmax {
		return 0
	}
	return rmin - lmax + 1
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
