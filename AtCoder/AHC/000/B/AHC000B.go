package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	d := nextInt()
	c, s, out := make([]int, 26), make([][]int, d), make([]int, d)
	for i := range c {
		c[i] = nextInt()
	}
	for i := range s {
		s[i] = make([]int, 26)
		for j := range s[i] {
			s[i][j] = nextInt()
		}
	}
	for i := range out {
		out[i] = nextInt() - 1
	}

	score(d, c, out, s)
}

func score(D int, c, out []int, s [][]int) int {
	score := 0
	last := make([]int, 26)

	for d := range out {
		last[out[d]] = d + 1
		for i := range last {
			score -= (d + 1 - last[i]) * c[i]
		}
		score += s[d][out[d]]
		fmt.Println(score)
	}

	return score
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
