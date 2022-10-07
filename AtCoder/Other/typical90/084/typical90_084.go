package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, s := nextInt(), nextLine()
	res := n * (n - 1) / 2

	last, l := s[0], 0
	for r := 1; r < len(s); r++ {
		if s[r] != last {
			res -= (l - r + 1) * (l - r) / 2
			last, l = s[r], r
		}
	}
	res -= ((len(s) - 1) - l + 1) * ((len(s) - 1) - l) / 2
	fmt.Println(res)
}

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
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

// var sc = bufio.NewScanner((os.Stdin))
