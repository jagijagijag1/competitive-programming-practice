package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K, A, max := nextInt(), nextInt(), []int{}, 0
	for i := 0; i < N; i++ {
		a := nextInt()
		A = append(A, a)
		if max < a {
			max = a
		}
	}

	if max < K {
		fmt.Println("IMPOSSIBLE")
		return
	}

	var g int
	if len(A) == 1 {
		g = A[0]
	} else if len(A) == 2 {
		g = gcd(A[0], A[1])
	} else {
		g = gcds(A[0], A[1], A[2:]...)
	}

	if K%g == 0 {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcds(a, b int, integers ...int) int {
	res := gcd(a, b)
	for i := 0; i < len(integers); i++ {
		res = gcd(res, integers[i])
	}
	return res
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
