package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, W := nextInt(), nextInt()
	maze := make([][]rune, H)
	for i := range maze {
		maze[i] = []rune(nextLine())
	}

	res := 0
	for sy := 0; sy < H; sy++ {
		for sx := 0; sx < W; sx++ {
			if maze[sy][sx] == '.' {
				res = max(res, bfs(sx, sy, H, W, maze))
			}
		}
	}
	fmt.Println(res)
}

type tuple struct {
	x, y int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func bfs(x, y, H, W int, maze [][]rune) int {
	dist, m := make([][]int, H), 0
	for i := range dist {
		dist[i] = make([]int, W)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	dist[y][x] = 0

	q := []tuple{}
	q = append(q, tuple{x, y})
	for len(q) > 0 {
		h := q[0]
		q = q[1:]

		for _, a := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			n := tuple{h.x + a[0], h.y + a[1]}

			if 0 <= n.x && n.x < W && 0 <= n.y && n.y < H && maze[n.y][n.x] == '.' && dist[n.y][n.x] == -1 {
				dist[n.y][n.x] = dist[h.y][h.x] + 1
				m = max(m, dist[n.y][n.x])
				q = append(q, n)
			}
		}

	}

	return m
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
