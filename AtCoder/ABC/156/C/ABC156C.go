package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC156C() {
	sc.Split(bufio.ScanWords)
	N, X := nextInt(), []int{}
	for i := 0; i < N; i++ {
		X = append(X, nextInt())
	}

	fmt.Printf("%d\n", ABC156C(N, X))
}

// ABC156C ...
func ABC156C(N int, X []int) (res int) {
	res = math.MaxInt32
	sort.Sort(sort.IntSlice(X))

	for i := X[0]; i <= X[len(X)-1]; i++ {
		tmp := 0
		for j := range X {
			tmp += (X[j] - i) * (X[j] - i)
		}
		res = minInt(res, tmp)
	}

	return
}

func minInt(a, b int) int {
	if a > b {
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
