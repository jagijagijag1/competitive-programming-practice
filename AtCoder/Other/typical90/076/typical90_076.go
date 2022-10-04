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
	a, sum := make([]int, 2*n), 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		sum += a[i]
	}
	for i := 0; i < n; i++ {
		a[n+i] = a[i]
	}

	r, s := 0, 0
	for l := 0; l < 2*n; l++ {
		for r < 2*n && s <= sum/10 {
			s += a[r]
			if s == sum/10 {
				fmt.Println("Yes")
				return
			}
			r++
		}
		if r == l {
			r++
		}
		s -= a[l]
	}
	fmt.Println("No")
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
