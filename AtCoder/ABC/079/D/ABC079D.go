package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, W, c := nextInt(), nextInt(), make([][]int, 10)
	for i := range c {
		c[i] = make([]int, 10)
		for j := range c[i] {
			c[i][j] = nextInt()
		}
	}
	A := make([][]int, H)
	for i := range A {
		A[i] = make([]int, W)
		for j := range A[i] {
			A[i][j] = nextInt()
		}
	}

	for k := 0; k < 10; k++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				c[i][j] = min(c[i][j], c[i][k]+c[k][j])
			}
		}
	}

	res := 0
	for i := range A {
		for j := range A[i] {
			if A[i][j] != -1 {
				res += c[A[i][j]][1]
			}
		}
	}
	fmt.Println(res)
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
