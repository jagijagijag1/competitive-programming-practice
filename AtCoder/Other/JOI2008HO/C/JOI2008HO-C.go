package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainJOI2008HOC() {
	sc.Split(bufio.ScanWords)
	N, M, P := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		P = append(P, nextInt())
	}

	fmt.Printf("%d\n", JOI2008HOC(M, P))
}

// JOI2008HOC ...
func JOI2008HOC(M int, P []int) (res uint) {
	Q := []int{}
	P = append(P, 0)
	for i := 0; i < len(P); i++ {
		for j := 0; j < len(P); j++ {
			Q = append(Q, P[i]+P[j])
		}
	}
	sort.Ints(Q)

	for i := 0; i < len(Q); i++ {
		idx := sort.Search(len(Q), func(j int) bool { return Q[i]+Q[j] > M })
		if idx > 0 {
			res = uint(maxInt(int(res), Q[i]+Q[idx-1]))
		}
	}

	return
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

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
