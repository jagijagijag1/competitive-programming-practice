package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, l, r, res := nextInt(), []float64{}, []float64{}, 0

	for i := 0; i < n; i++ {
		t, ll, rr := nextInt(), float64(nextInt()), float64(nextInt())
		if t == 2 {
			rr -= 0.5
		} else if t == 3 {
			ll += 0.5
		} else if t == 4 {
			ll += 0.5
			rr -= 0.5
		}
		l, r = append(l, ll), append(r, rr)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if max(l[i], l[j]) <= min(r[i], r[j]) {
				res++
			}
		}
	}
	fmt.Println(res)
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b float64) float64 {
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
