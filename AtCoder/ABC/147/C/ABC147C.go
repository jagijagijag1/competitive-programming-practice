package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	x, y := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		a := nextInt()
		x[i], y[i] = make([]int, a), make([]int, a)
		for j := 0; j < a; j++ {
			x[i][j] = nextInt() - 1
			y[i][j] = nextInt()
		}
	}

	res := 0
	for bit := 0; bit < (1 << uint(n)); bit++ {
		tmp, flag := 0, true
		for i := 0; i < n; i++ {
			if bit>>uint(i)&1 == 1 {
				tmp++
				for j := 0; j < len(x[i]); j++ {
					if bit>>uint(x[i][j])&1 != y[i][j] {
						flag = false
						break
					}
				}
				if flag == false {
					break
				}
			}
		}
		if flag {
			res = max(res, tmp)
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
