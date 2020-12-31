package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC160C() {
	sc.Split(bufio.ScanWords)
	K, N, A := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	fmt.Printf("%d\n", ABC160C(K, A))
}

// ABC160C ...
func ABC160C(K int, A []int) (res int) {
	dists, last := []int{}, 0
	for i := 0; i < len(A); i++ {
		dists = append(dists, A[i]-last)
		last = A[i]
	}
	dists[0] += K - last

	max := -1
	for i := 0; i < len(dists); i++ {
		res += dists[i]
		max = maxInt(max, dists[i])
	}

	return res - max
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
