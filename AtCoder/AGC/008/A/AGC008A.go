package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	x, y := nextInt(), nextInt()

	var res int
	if 0 <= x && 0 <= y {
		if x <= y {
			res = y - x
		} else {
			res = x - y + 2
		}
	} else if x < 0 && y < 0 {
		if abs(x) < abs(y) {
			res = abs(y) - abs(x) + 2
		} else {
			res = abs(x) - abs(y)
		}
	} else {
		res = max(abs(x), abs(y)) - min(abs(x), abs(y)) + 1
	}
	if y == 0 {
		res--
	}

	fmt.Println(res)
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
