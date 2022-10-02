package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, q := nextInt(), nextInt()
	a, d, res := make([]int, n), make([]int, n-1), 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		if i != 0 {
			d[i-1] = a[i] - a[i-1]
			res += abs(d[i-1])
		}
	}

	for i := 0; i < q; i++ {
		l, r, v := nextInt()-1, nextInt()-1, nextInt()
		if 0 < l {
			pre := abs(d[l-1])
			d[l-1] += v
			res += abs(d[l-1]) - pre
		}
		if r < n-1 {
			pre := abs(d[r])
			d[r] -= v
			res += abs(d[r]) - pre
		}
		fmt.Println(res)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
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
