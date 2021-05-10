package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, x, y, h := nextInt(), []int{}, []int{}, []int{}
	for i := 0; i < N; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
		h = append(h, nextInt())
	}

	for yy := 0; yy <= 100; yy++ {
		for xx := 0; xx <= 100; xx++ {
			th := -1
			for i := 0; i < N; i++ {
				if h[i] > 0 {
					t := h[i] + abs(yy-y[i]) + abs(xx-x[i])
					if th == -1 {
						th = t
					} else if th != t {
						th = -2
						break
					}
				}
			}
			if th == -2 {
				continue
			}

			for i := 0; i < N; i++ {
				if h[i] == 0 {
					d := abs(yy-y[i]) + abs(xx-x[i])
					if d < th {
						th = -2
						break
					}
				}
			}
			if th == -2 {
				continue
			}

			fmt.Println(xx, yy, th)
			return
		}
	}
}

func abs(a int) int {
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
