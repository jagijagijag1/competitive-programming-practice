package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainABC119C() {
	sc.Split(bufio.ScanWords)
	N, A, B, C, l := nextInt(), nextInt(), nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		l = append(l, nextInt())
	}

	fmt.Printf("%d\n", ABC119C(A, B, C, l))
}

// ABC119C ...
func ABC119C(A, B, C int, l []int) int {
	return dfsABC119C(l, A, B, C, 0, 0, 0, 0)
}

func dfsABC119C(l []int, A, B, C, a, b, c, idx int) int {
	if idx == len(l) {
		if a > 0 && b > 0 && c > 0 {
			return absInt(A-a) + absInt(B-b) + absInt(C-c) - 30
		}
		return math.MaxInt32
	}

	r0 := dfsABC119C(l, A, B, C, a, b, c, idx+1)
	r1 := dfsABC119C(l, A, B, C, a+l[idx], b, c, idx+1) + 10
	r2 := dfsABC119C(l, A, B, C, a, b+l[idx], c, idx+1) + 10
	r3 := dfsABC119C(l, A, B, C, a, b, c+l[idx], idx+1) + 10

	return minIntInSlice([]int{r0, r1, r2, r3})
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minIntInSlice(v []int) (res int) {
	res = math.MaxInt32
	for i, e := range v {
		if i == 0 || res > e {
			res = e
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
