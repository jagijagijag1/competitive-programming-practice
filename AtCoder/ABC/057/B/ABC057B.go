package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainABC057B() {
	sc.Split(bufio.ScanWords)
	N, M, a, b, c, d := nextInt(), nextInt(), []int{}, []int{}, []int{}, []int{}
	for i := 0; i < N; i++ {
		a, b = append(a, nextInt()), append(b, nextInt())
	}
	for j := 0; j < M; j++ {
		c, d = append(c, nextInt()), append(d, nextInt())
	}

	for _, r := range ABC057B(a, b, c, d) {
		fmt.Printf("%d\n", r)
	}
	// fmt.Printf("%d\n", ABC057B())
}

// ABC057B ...
func ABC057B(a, b, c, d []int) (res []int) {
	for i := range a {
		mind, minidx := math.MaxInt32, math.MaxInt32
		for j := range c {
			d := absInt(a[i]-c[j]) + absInt(b[i]-d[j])
			if d < mind {
				mind = d
				minidx = j
			}
		}
		res = append(res, minidx+1)
	}

	return
}

func absInt(a int) int {
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
