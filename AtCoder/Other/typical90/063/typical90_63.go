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
	p := make([][]int, h)
	for i := range p {
		p[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			p[i][j] = nextInt()
		}
	}

	res := 0
	for bit := 0; bit < (1 << uint(h)); bit++ {
		m, mmax := map[int]int{}, 0
		for ww := 0; ww < w; ww++ {
			cur, isSameNum := -1, true
			for hh := 0; hh < h; hh++ {
				if bit>>uint(hh)&1 == 1 {
					if cur == -1 {
						cur = p[hh][ww]
					} else {
						if cur != p[hh][ww] {
							isSameNum = false
							break
						}
					}
				}
			}
			if isSameNum {
				m[cur]++
				mmax = max(mmax, m[cur])
			}
		}

		hcnt := 0
		for hh := 0; hh < h; hh++ {
			if bit>>uint(hh)&1 == 1 {
				hcnt++
			}
		}
		res = max(res, hcnt*mmax)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
		return b
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
