package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	s, res := nextLine(), 0
	l, r := 0, len(s)-1

	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			if s[l] == 'x' {
				res++
				l++
			} else if s[r] == 'x' {
				res++
				r--
			} else {
				fmt.Println(-1)
				return
			}
		}
	}
	fmt.Println(res)
}

// var sc = bufio.NewScanner((os.Stdin))

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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)
