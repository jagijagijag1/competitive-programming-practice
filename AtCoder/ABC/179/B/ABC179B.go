package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, f1, f2 := nextInt(), false, false
	if nextInt() == nextInt() {
		f2 = true
	}
	if nextInt() == nextInt() {
		f1 = true
	}

	for i := 0; i < n-2; i++ {
		d1, d2, f := nextInt(), nextInt(), false
		if d1 == d2 {
			f = true
		}

		if f2 && f1 && f {
			fmt.Println("Yes")
			return
		}
		f2, f1 = f1, f
	}
	fmt.Println("No")
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
