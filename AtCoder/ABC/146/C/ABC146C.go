package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	A, B, X := nextInt(), nextInt(), nextInt()

	if X < A+B {
		fmt.Println(0)
		return
	}

	minN, maxN := 0, 1000000000
	maxcost := A*maxN + B*d(maxN)
	if maxcost < X {
		fmt.Println(maxN)
		return
	}

	for minN < maxN-1 {
		mid := (minN + maxN) / 2
		cost := A*mid + B*d(mid)
		if cost <= X {
			minN = mid
		} else {
			maxN = mid
		}
	}
	fmt.Println(minN)
}

func d(N int) int {
	return len(strconv.Itoa(N))
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
