package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	c := make([]int, N)
	for i := 0; i < N; i++ {
		a := nextInt()
		idx := (N - a - 1) / 2
		c[idx]++
		if N%2 == 1 && idx == (N-1)/2 && 1 < c[idx] {
			fmt.Println(0)
			return
		} else if 2 < c[idx] {
			fmt.Println(0)
			return
		}
	}

	res := 1
	for i := 0; i < N/2; i++ {
		res *= 2
		res %= 1000000007
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
