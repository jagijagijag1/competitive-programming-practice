package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, l, li, r, ri := nextInt(), 0, 0, 0, 0
	nn := pow(2, n)

	for i := 0; i < nn; i++ {
		a := nextInt()
		if i < nn/2 {
			if l < a {
				l, li = a, i+1
			}
		} else {
			if r < a {
				r, ri = a, i+1
			}
		}
	}

	if l < r {
		fmt.Println(li)
	} else {
		fmt.Println(ri)
	}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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
