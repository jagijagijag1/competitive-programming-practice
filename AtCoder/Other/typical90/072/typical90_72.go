package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var h, w int
var dx, dy []int
var used [][]bool
var board []string

func main() {
	sc.Split(bufio.ScanWords)
	h, w = nextInt(), nextInt()
	board = make([]string, h)
	for i := 0; i < h; i++ {
		board[i] = nextLine()
	}

	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
	used = make([][]bool, h)
	for i := 0; i < h; i++ {
		used[i] = make([]bool, w)
	}

	res := -1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			res = max(res, dfs(i, j, i, j))
		}
	}

	if res <= 2 {
		res = -1
	}
	fmt.Println(res)
}

func dfs(sx, sy, px, py int) int {
	if sx == px && sy == py && used[px][py] == true {
		return 0
	}

	used[px][py] = true
	res := math.MinInt64
	for i := 0; i < 4; i++ {
		nx, ny := px+dx[i], py+dy[i]
		if nx < 0 || h <= nx || ny < 0 || w <= ny {
			continue
		}
		if board[nx][ny] == '#' {
			continue
		}
		if used[nx][ny] == true && !(nx == sx && ny == sy) {
			continue
		}

		res = max(res, dfs(sx, sy, nx, ny)+1)
	}
	used[px][py] = false

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
