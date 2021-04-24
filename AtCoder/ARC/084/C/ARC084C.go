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
	N, A, B, C := nextInt(), []int{}, []int{}, []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}
	for i := 0; i < N; i++ {
		B = append(B, nextInt())
	}
	for i := 0; i < N; i++ {
		C = append(C, nextInt())
	}
	sort.Ints(A)
	sort.Ints(B)
	sort.Ints(C)

	res := 0
	for i := 0; i < N; i++ {
		ai := sort.Search(len(A), func(j int) bool {
			return A[j] >= B[i]
		})
		ci := sort.Search(len(C), func(j int) bool {
			return C[j] > B[i]
		})
		res += ai * (N - ci)
	}
	fmt.Println(res)
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
