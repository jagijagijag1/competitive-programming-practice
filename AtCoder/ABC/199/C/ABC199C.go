package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, s, q, f := nextInt(), []rune(nextLine()), nextInt(), false
	for i := 0; i < q; i++ {
		t, a, b := nextInt(), nextInt()-1, nextInt()-1
		if t == 1 {
			if f {
				a, b = (a+n)%(2*n), (b+n)%(2*n)
			}
			s[a], s[b] = s[b], s[a]
		} else {
			f = !f
		}
	}
	if f {
		ss := string(s)
		fmt.Println(ss[n:] + ss[:n])
	} else {
		fmt.Println(string(s))
	}
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
