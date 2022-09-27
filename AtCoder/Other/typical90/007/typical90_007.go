package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)

	for i := range a {
		a[i] = nextInt()
	}
	sort.Ints(a)

	q := nextInt()
	for i := 0; i < q; i++ {
		b := nextInt()
		idx := sort.Search(len(a), func(i int) bool {
			return a[i] >= b
		})

		b1, b2 := abs(a[max(0, idx-1)]-b), abs(a[min(idx, n-1)]-b)
		if b1 < b2 {
			fmt.Println(b1)
		} else {
			fmt.Println(b2)
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
