package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()
	a, diff := make([][]int, h), make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
		}
	}
	for i := 0; i < h; i++ {
		diff[i] = make([]int, w)
		for j := 0; j < w; j++ {
			diff[i][j] = a[i][j] - nextInt()
		}
	}

	res := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			if diff[i][j] != 0 {
				diff[i+1][j] -= diff[i][j]
				diff[i][j+1] -= diff[i][j]
				diff[i+1][j+1] -= diff[i][j]
				res += abs(diff[i][j])
				diff[i][j] = 0
			}
		}
	}

	for i := 0; i < h; i++ {
		if diff[i][w-1] != 0 {
			fmt.Println("No")
			return
		}
	}
	for j := 0; j < w; j++ {
		if diff[h-1][j] != 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	fmt.Println(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
