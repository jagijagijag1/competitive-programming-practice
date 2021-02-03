package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, B := nextInt(), []int{}, []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
		B = append(B, nextInt())
	}

	res := 0
	for i := N - 1; 0 <= i; i-- {
		a, b := A[i]+res, B[i]
		if b < a {
			if a%b == 0 {
				b += b * (a - b) / b
			} else {
				b += b * ((a-b)/b + 1)
			}
		}
		res += b - a
	}
	// for i := N - 1; 0 <= i; i-- {
	// 	a := A[i] + res
	// 	m := a % B[i]
	// 	if m != 0 {
	// 		res += B[i] - m
	// 	}
	// }

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
