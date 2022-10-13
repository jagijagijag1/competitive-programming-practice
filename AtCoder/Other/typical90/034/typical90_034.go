package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	res, m, r := 0, map[int]int{}, 0
	for l := 0; l < n; l++ {
		for r < n {
			m[a[r]]++
			if k < len(m) {
				m[a[r]]--
				break
			}
			r++
		}

		res = max(res, r-l)
		m[a[l]]--
		if m[a[l]] == 0 {
			delete(m, a[l])
		}

		if r == l {
			r++
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
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
