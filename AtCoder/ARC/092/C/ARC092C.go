package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainARC092C() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	a, b, c, d := []int{}, []int{}, []int{}, []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}
	for i := 0; i < N; i++ {
		c = append(c, nextInt())
		d = append(d, nextInt())
	}

	fmt.Printf("%d\n", ARC092C(a, b, c, d))
}

// ARC092C ...
func ARC092C(a, b, c, d []int) int {
	R, B, used, pair := []Point{}, []Point{}, make([]bool, len(a)), map[Point]Point{}
	for i := 0; i < len(a); i++ {
		R = append(R, Point{a[i], b[i]})
		B = append(B, Point{c[i], d[i]})
	}
	sort.Slice(R, func(i, j int) bool {
		if R[i].y > R[j].y {
			return true
		} else if R[i].y == R[j].y {
			if R[i].x > R[j].x {
				return true
			}
		}
		return false
	})
	sort.Slice(B, func(i, j int) bool {
		if B[i].x < B[j].x {
			return true
		} else if B[i].x == B[j].x {
			if B[i].y < B[j].y {
				return true
			}
		}
		return false
	})

	for bi := 0; bi < len(B); bi++ {
		for ri := 0; ri < len(R); ri++ {
			if used[ri] {
				continue
			}

			if B[bi].x > R[ri].x && B[bi].y > R[ri].y {
				pair[B[bi]] = R[ri]
				used[ri] = true
				break
			}
		}
	}

	return len(pair)
}

// Point ...
type Point struct {
	x, y int
}

func pointdist(p1, p2 Point) float64 {
	xd := p1.x*p1.x + p2.x*p2.x
	yd := p1.y*p1.y + p2.y*p2.y

	return math.Sqrt(float64(xd + yd))
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

// ARC092Cdinic ...
func ARC092Cdinic(a, b, c, d []int) int {
	N, R, B := len(a), []Point{}, []Point{}
	G, level, lastCheckedEdgeIndex = make([][]*DinicEdge, 2*N+2), make([]int, 2*N+2), make([]int, 2*N+2)

	for i := 0; i < len(a); i++ {
		R = append(R, Point{a[i], b[i]})
		B = append(B, Point{c[i], d[i]})
	}

	// s(source): 0, red: 1 - N, blue: N+1 - 2N, t(sink): 2N+1
	for ri := 0; ri < N; ri++ {
		for bi := 0; bi < N; bi++ {
			if R[ri].x < B[bi].x && R[ri].y < B[bi].y {
				addDinicEdge(ri+1, bi+N+1, 1)
			}
		}
	}
	for i := 0; i < N; i++ {
		addDinicEdge(0, i+1, 1)
		addDinicEdge(i+N+1, 2*N+1, 1)
	}

	return maxFlow(0, 2*N+1)
}

// DinicEdge ...
type DinicEdge struct {
	to, cap, rev int
}

// G is graph
var G [][]*DinicEdge
var level []int
var lastCheckedEdgeIndex []int

func addDinicEdge(from, to, cap int) {
	G[from] = append(G[from], &DinicEdge{to, cap, len(G[to])})
	G[to] = append(G[to], &DinicEdge{from, 0, len(G[from]) - 1})
}

func bfs(s int) {
	for i := 0; i < len(level); i++ {
		level[i] = -1
	}

	level[s] = 0
	queue := []int{s}
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		for i := 0; i < len(G[v]); i++ {
			e := G[v][i]
			if e.cap > 0 && level[e.to] < 0 {
				level[e.to] = level[v] + 1
				queue = append(queue, e.to)
			}
		}
	}
}

func dfs(v, t, f int) int {
	if v == t {
		return f
	}

	for i := lastCheckedEdgeIndex[v]; i < len(G[v]); i++ {
		e := G[v][i]
		if e.cap > 0 && level[v] < level[e.to] {
			d := dfs(e.to, t, minInt(f, e.cap))
			if d > 0 {
				e.cap -= d
				G[e.to][e.rev].cap += d
				return d
			}
		}
	}

	return 0
}

func maxFlow(s, t int) (flow int) {
	for {
		bfs(s)
		if level[t] < 0 {
			return
		}

		for i := 0; i < len(lastCheckedEdgeIndex); i++ {
			lastCheckedEdgeIndex[i] = 0
		}

		f := dfs(s, t, 100000000)
		for f > 0 {
			flow += f
			f = dfs(s, t, 100000000)
		}
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
