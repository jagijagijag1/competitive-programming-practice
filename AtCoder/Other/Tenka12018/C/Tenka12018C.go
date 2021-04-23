package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A := nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	res := 0
	if N%2 == 0 {
		for i := 0; i < N/2-1; i++ {
			res += A[i] * 2
		}
		res += A[N/2-1]
		res -= A[N/2]
		for i := N/2 + 1; i < N; i++ {
			res -= A[i] * 2
		}
	} else {
		t1, t2 := 0, 0
		for i := 0; i < N/2-1; i++ {
			t1 += A[i] * 2
		}
		t1 += A[N/2-1] + A[N/2]
		for i := N/2 + 1; i < N; i++ {
			t1 -= A[i] * 2
		}

		for i := 0; i < N/2; i++ {
			t2 += A[i] * 2
		}
		t2 -= A[N/2] + A[N/2+1]
		for i := N/2 + 2; i < N; i++ {
			t2 -= A[i] * 2
		}
		res = max(t1, t2)
	}
	fmt.Println(res)
}

func max(a, b int) int {
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

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
