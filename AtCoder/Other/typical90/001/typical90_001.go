package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, l, k := nextInt(), nextInt(), nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	left, right := 0, l
	for left < right-1 {
		mid := (left + right) / 2
		last, num := 0, 0
		for i := 0; i < n; i++ {
			if mid <= a[i]-last && mid <= l-a[i] {
				last = a[i]
				num++
			}
		}
		if k <= num {
			left = mid
		} else {
			right = mid
		}
	}
	fmt.Println(left)
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
