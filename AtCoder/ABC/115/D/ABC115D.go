package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var a, p []int

func main() {
	sc.Split(bufio.ScanWords)
	N, X := nextInt(), nextInt()
	a, p = append(a, 1), append(p, 1)
	for i := 0; i < N-1; i++ {
		a = append(a, a[i]*2+3)
		p = append(p, p[i]*2+1)
	}

	fmt.Println(f(N, X))
}

func f(N, X int) int {
	if N == 0 {
		if X <= 0 {
			return 0
		} else {
			return 1
		}
	} else if X <= 1+a[N-1] {
		return f(N-1, X-1)
	} else {
		return p[N-1] + 1 + f(N-1, X-2-a[N-1])
	}
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
