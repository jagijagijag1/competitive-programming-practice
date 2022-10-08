package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var G [][]int
var dist []int

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	G, dist = make([][]int, n), make([]int, n)
	for i := range G {
		G[i] = []int{}
	}
	for i := 0; i < n-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	bfs(0)
	maxdist, maxnode := -1, -1
	for i := 0; i < n; i++ {
		if maxdist < dist[i] {
			maxdist = dist[i]
			maxnode = i
		}
	}

	bfs(maxnode)
	res := -1
	for i := 0; i < n; i++ {
		res = max(res, dist[i])
	}
	fmt.Println(res + 1)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func bfs(v int) {
	for i := range dist {
		dist[i] = math.MaxInt64
	}

	q := []int{v}
	dist[v] = 0
	for 0 < len(q) {
		from := q[0]
		q = q[1:]

		for _, to := range G[from] {
			if dist[to] == math.MaxInt64 {
				dist[to] = dist[from] + 1
				q = append(q, to)
			}
		}
	}

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
	initialBufSize = 100000
	maxBufSize     = 10000000
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
