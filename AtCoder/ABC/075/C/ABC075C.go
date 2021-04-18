package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()
	adj := make([][]int, N)

	for i := 0; i < M; i++ {
		s, t := nextInt()-1, nextInt()-1
		adj[s] = append(adj[s], t)
		adj[t] = append(adj[t], s)
	}

	bridges(adj, N)

	fmt.Println(len(res))
}

var res []edge
var cnt int

func bridges(G [][]int, V int) {
	low, pre := make([]int, V), make([]int, V)
	for i := range low {
		low[i], pre[i] = -1, -1
	}

	dfs(G, 0, -1, low, pre)
}

func dfs(g [][]int, v, from int, low, pre []int) int {
	pre[v] = cnt
	low[v] = pre[v]
	cnt++
	for _, to := range g[v] {
		if pre[to] == -1 {
			low[v] = min(low[v], dfs(g, to, v, low, pre))
			if low[to] == pre[to] {
				res = append(res, edge{v, to})
			}
		} else {
			if from == to {
				continue
			}
			low[v] = min(low[v], low[to])
		}
	}
	return low[v]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type edge struct {
	s, t int
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
