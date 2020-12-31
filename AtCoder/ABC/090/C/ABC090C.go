package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

// func main() {
func mainABC090C() {
	sc.Split(bufio.ScanWords)
	N, A := nextInt(), [][]int{{}, {}}

	for i := 0; i < 2; i++ {
		for j := 0; j < N; j++ {
			A[i] = append(A[i], nextInt())
		}
	}

	fmt.Printf("%d\n", ABC090C(A))
}

// ABC090C ...
func ABC090C(A [][]int) int {
	res := 0

	sum2 := 0
	for i := 0; i < len(A[1]); i++ {
		sum2 += A[1][i]
	}

	sum1tmp, sum2tmp := 0, sum2
	for i := 0; i < len(A[0]); i++ {
		sum1tmp += A[0][i]
		sumtmp := sum1tmp + sum2tmp
		if res < sumtmp {
			res = sumtmp
		}
		sum2tmp -= A[1][i]
	}

	return res
}
