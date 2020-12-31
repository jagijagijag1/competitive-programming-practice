package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC074B() {
	sc.Split(bufio.ScanWords)
	N, K, x := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		x = append(x, nextInt())
	}

	fmt.Printf("%d\n", ABC074B(K, x))
}

// ABC074B ...
func ABC074B(K int, x []int) (res int) {
	for _, xi := range x {
		t := absInt(xi - K)
		if t < xi {
			res += 2 * t
		} else {
			res += 2 * xi
		}
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
