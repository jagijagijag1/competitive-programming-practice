package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, a, b, c := nextInt(), []int{}, []int{}, []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}
	for i := 0; i < N; i++ {
		b = append(b, nextInt())
	}
	for i := 0; i < N; i++ {
		c = append(c, nextInt())
	}

	d := make([]int, N+1)
	for i := 0; i < N; i++ {
		d[a[i]]++
	}

	res := 0
	for j := 0; j < N; j++ {
		res += d[b[c[j]-1]]
	}
	fmt.Println(res)
}

// var sc = bufio.NewScanner((os.Stdin))

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
