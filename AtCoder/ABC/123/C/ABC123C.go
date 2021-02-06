package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, B, C, D, E := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	counts := []int{ciel(N, A), ciel(N, B), ciel(N, C), ciel(N, D), ciel(N, E)}
	m := maxIntInSlice(counts)
	fmt.Println(m + 4)
}

func ciel(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}

func maxIntInSlice(v []int) (res int) {
	res = math.MinInt32
	for i, e := range v {
		if i == 0 || res < e {
			res = e
		}
	}
	return
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
