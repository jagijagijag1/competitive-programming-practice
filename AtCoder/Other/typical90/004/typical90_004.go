package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()
	a, res := make([][]int, h), make([][]int, h)
	for i := range a {
		a[i] = make([]int, w)
		res[i] = make([]int, w)
	}

	sh, sw := make([]int, h), make([]int, w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
			sh[i] += a[i][j]
			sw[j] += a[i][j]
		}
	}

	for i := range a {
		for j := range a[i] {
			res[i][j] = sh[i] + sw[j] - a[i][j]
		}
		fmt.Println(strings.Trim(fmt.Sprint(res[i]), "[]"))
	}
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
