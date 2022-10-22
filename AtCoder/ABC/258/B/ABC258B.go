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
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		aa := nextLine()
		for j := 0; j < n; j++ {
			a[i][j] = int(aa[j] - '0')
		}
	}

	res := 0
	dx := []int{1, 1, 1, 0, 0, -1, -1, -1}
	dy := []int{1, 0, -1, 1, -1, 1, 0, -1}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < 8; k++ {
				tmp := 0
				cx, cy := i, j
				for l := 0; l < n; l++ {
					tmp *= 10
					tmp += a[cx][cy]
					cx, cy = (cx+dx[k]+n)%n, (cy+dy[k]+n)%n
				}
				res = max(res, tmp)
			}
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
