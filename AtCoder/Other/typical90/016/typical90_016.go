package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a, b, c := nextInt(), nextInt(), nextInt()

	maxcoin, res := 10000, 10000
	for i := 0; i < maxcoin; i++ {
		for j := 0; j < (maxcoin - i); j++ {
			sum := a*i + b*j
			if n < sum {
				break
			}
			if (n-sum)%c == 0 {
				tmp := i + j + ((n - sum) / c)
				res = min(res, tmp)
			}
		}
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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

// var sc = bufio.NewScanner((os.Stdin))
