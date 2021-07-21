package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, r, b := nextInt(), 0, 0
	for i := 0; i < n; i++ {
		s := nextLine()
		for _, ss := range s {
			if ss == 'R' {
				r++
			} else if ss == 'B' {
				b++
			}
		}
	}

	if b < r {
		fmt.Println("TAKAHASHI")
	} else if r < b {
		fmt.Println("AOKI")
	} else {
		fmt.Println("DRAW")
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
