package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	A, B, C, X, Y := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	base := A*X + B*Y
	comb := max(X, Y) * C * 2

	var mincomb int
	if X > Y {
		mincomb = Y*C*2 + A*(X-Y)
	} else {
		mincomb = X*C*2 + B*(Y-X)
	}

	fmt.Println(min(base, min(comb, mincomb)))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
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
