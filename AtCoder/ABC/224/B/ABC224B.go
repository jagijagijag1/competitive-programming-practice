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
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			a[i][j] = nextInt()
		}
	}

	for i1 := 0; i1 < h; i1++ {
		for i2 := i1 + 1; i2 < h; i2++ {
			for j1 := 0; j1 < w; j1++ {
				for j2 := j1 + 1; j2 < w; j2++ {
					if a[i1][j1]+a[i2][j2] > a[i2][j1]+a[i1][j2] {
						fmt.Println("No")
						return
					}
				}
			}
		}
	}
	fmt.Println("Yes")
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

// var sc = bufio.NewScanner((os.Stdin))
