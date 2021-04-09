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
	x, y, a, b, c := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	p, q, r := []int{}, []int{}, []int{}
	for i := 0; i < a; i++ {
		p = append(p, nextInt())
	}
	for i := 0; i < b; i++ {
		q = append(q, nextInt())
	}
	for i := 0; i < c; i++ {
		r = append(r, nextInt())
	}
	sort.Ints(p)
	sort.Ints(q)

	p = p[len(p)-x:]
	q = q[len(q)-y:]
	r = append(r, p...)
	r = append(r, q...)
	sort.Ints(r)

	res := 0
	for i := 0; i < x+y; i++ {
		res += r[len(r)-i-1]
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
