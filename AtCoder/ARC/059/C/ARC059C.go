package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, a, sum := nextInt(), []int{}, 0
	for i := 0; i < N; i++ {
		aa := nextInt()
		a = append(a, aa)
		sum += aa
	}

	m, res1, res2 := sum/N, 0, 0
	for _, aa := range a {
		res1 += abs(aa-m) * abs(aa-m)
		res2 += abs(aa-(m+1)) * abs(aa-(m+1))
	}

	fmt.Println(min(res1, res2))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
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
