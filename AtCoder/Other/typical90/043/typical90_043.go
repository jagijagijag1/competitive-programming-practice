package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()
	rs, cs, rt, ct := nextInt()-1, nextInt()-1, nextInt()-1, nextInt()-1
	board := make([]string, h)
	for i := 0; i < h; i++ {
		board[i] = nextLine()
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	dist := make([][][4]int, h)
	for i := range dist {
		dist[i] = make([][4]int, w)
		for j := range dist[i] {
			dist[i][j][0] = math.MaxInt64
			dist[i][j][1] = math.MaxInt64
			dist[i][j][2] = math.MaxInt64
			dist[i][j][3] = math.MaxInt64
		}
	}

	// 01-bfs
	deq := Deque()
	for i := 0; i < 4; i++ {
		dist[rs][cs][i] = 0
		deq.PushBack(state{rs, cs, i})
	}
	for deq.Len() != 0 {
		u := deq.GetFront()
		deq.RemoveFront()
		for i := 0; i < 4; i++ {
			nx, ny := u.x+dx[i], u.y+dy[i]
			cost := dist[u.x][u.y][u.dir]
			if u.dir != i {
				cost++
			}

			if 0 <= nx && nx < h && 0 <= ny && ny < w {
				if board[nx][ny] == '.' && cost < dist[nx][ny][i] {
					dist[nx][ny][i] = cost
					if u.dir != i {
						deq.PushBack(state{nx, ny, i})
					} else {
						deq.PushFront(state{nx, ny, i})
					}
				}
			}
		}

	}

	res := math.MaxInt64
	for i := 0; i < 4; i++ {
		res = min(res, dist[rt][ct][i])
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type state struct {
	x, y, dir int
}

/*** deque ***/ // dq := Deque()
type deque_type struct {
	Dat []state
	L   int
	R   int
}

func Deque() *deque_type {
	return &deque_type{Dat: make([]state, 20000001), L: 10000000, R: 10000001}
}
func (dq *deque_type) PushBack(val state) {
	dq.Dat[dq.R] = val
	dq.R++
}
func (dq *deque_type) PushFront(val state) {
	dq.Dat[dq.L] = val
	dq.L--
}
func (dq *deque_type) RemoveBack() {
	dq.R--
}
func (dq *deque_type) RemoveFront() {
	dq.L++
}
func (dq *deque_type) GetBack() state {
	return dq.Dat[dq.R-1]
}
func (dq *deque_type) GetFront() state {
	return dq.Dat[dq.L+1]
}
func (dq *deque_type) GetIth(idx int) state {
	return dq.Dat[dq.L+idx+1]
}
func (dq *deque_type) Len() int {
	return dq.R - dq.L - 1
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
