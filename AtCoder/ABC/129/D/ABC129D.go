package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, W, s := nextInt(), nextInt(), []string{}

	for i := 0; i < H; i++ {
		s = append(s, nextLine())
	}

	l, r, u, d := make([][]int, H), make([][]int, H), make([][]int, H), make([][]int, H)
	for i := range l {
		l[i], r[i], u[i], d[i] = make([]int, W), make([]int, W), make([]int, W), make([]int, W)
	}

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if s[h][w] == '.' {
				if w == 0 {
					l[h][w] = 1
					continue
				}
				l[h][w] = l[h][w-1] + 1
			}
		}
	}
	for h := H - 1; 0 <= h; h-- {
		for w := W - 1; 0 <= w; w-- {
			if s[h][w] == '.' {
				if w == W-1 {
					r[h][w] = 1
					continue
				}
				r[h][w] = r[h][w+1] + 1
			}
		}
	}
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if s[h][w] == '.' {
				if h == 0 {
					u[h][w] = 1
					continue
				}
				u[h][w] = u[h-1][w] + 1
			}
		}
	}
	for h := H - 1; 0 <= h; h-- {
		for w := W - 1; 0 <= w; w-- {
			if s[h][w] == '.' {
				if h == H-1 {
					d[h][w] = 1
					continue
				}
				d[h][w] = d[h+1][w] + 1
			}
		}
	}

	res := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			res = max(res, l[h][w]+r[h][w]+u[h][w]+d[h][w]-3)
		}
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
		return b
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
