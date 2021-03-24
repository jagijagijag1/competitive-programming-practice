package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, s := nextInt(), []int{}, 0
	for i := 0; i < N; i++ {
		a := nextInt()
		A = append(A, a)
		s += a
	}
	for i := 1; i < N; i += 2 {
		s -= 2 * A[i]
	}

	res := []int{s}
	for i := 1; i < N; i++ {
		res = append(res, 2*A[i-1]-res[i-1])
	}
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
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
