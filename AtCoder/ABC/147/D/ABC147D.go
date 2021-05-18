package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, mod := nextInt(), []int{}, 1000000007
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	res := 0
	for i := 0; i < 60; i++ {
		x := 0
		for j := 0; j < N; j++ {
			if A[j]>>i&1 == 1 {
				x++
			}
		}

		now := x * (N - x) % mod
		for j := 0; j < i; j++ {
			now = (now * 2) % mod
		}

		res += now
		res %= mod
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
