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
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	sort.Ints(a)
	sort.Ints(b)

	fmt.Println(dist(a[n/2], b[n/2], a, b))
}

func dist(s, g int, a, b []int) (res int) {
	for i := 0; i < len(a); i++ {
		res += abs(s - a[i])
		res += abs(a[i] - b[i])
		res += abs(g - b[i])
	}
	return
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
