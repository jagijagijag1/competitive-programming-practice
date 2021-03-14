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
	n, a := nextInt(), []int{}
	for i := 0; i < n; i++ {
		a = append(a, nextInt())
	}
	sort.Ints(a)

	min, maxn, maxr := 1000000000, a[n-1], 0
	for i := n - 2; 0 <= i; i-- {
		r := abs(maxn/2 - a[i])
		if r < min {
			min, maxr = r, a[i]
		}
	}

	fmt.Println(maxn, maxr)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
