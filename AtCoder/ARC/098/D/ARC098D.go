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

//func main() {
func mainARC098D() {
	sc.Split(bufio.ScanWords)
	N, A := nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	fmt.Printf("%d\n", ARC098D(A))
}

// ARC098D ...
func ARC098D(A []int) int {
	//res := len(A)
	res := 0
	r, sum := 0, 0
	for l := 0; l < len(A); l++ {
		for r < len(A) && sum^A[r] == sum+A[r] {
			sum += A[r]
			r++
		}

		res += r - l
		if r == l {
			r++
		} else {
			sum -= A[l]
		}
	}

	return res
}
