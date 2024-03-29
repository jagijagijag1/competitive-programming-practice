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
	a, b, c := make([]int, 46), make([]int, 46), make([]int, 46)
	for i := 0; i < n; i++ {
		a[nextInt()%46]++
	}
	for i := 0; i < n; i++ {
		b[nextInt()%46]++
	}
	for i := 0; i < n; i++ {
		c[nextInt()%46]++
	}

	res := 0
	for i := range a {
		for j := range b {
			for k := range c {
				if (i+j+k)%46 == 0 {
					res += a[i] * b[j] * c[k]
				}
			}
		}
	}
	fmt.Println(res)
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
