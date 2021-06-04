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
	s := make([]int, n)
	s[0] = a[0]
	for i := 1; i < n; i++ {
		s[i] = a[i] + s[i-1]
	}

	res := 0
	for i := 1; i < n; i++ {
		res += a[i]*i - s[i-1]
	}
	fmt.Println(res)
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
