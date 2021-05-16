package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var G [][]edge
var distFromK []int

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	G = make([][]edge, N)
	for i := range G {
		G[i] = []edge{}
	}
	for i := 0; i < N-1; i++ {
		a, b, c := nextInt()-1, nextInt()-1, nextInt()
		G[a] = append(G[a], edge{b, c})
		G[b] = append(G[b], edge{a, c})
	}

	Q, K := nextInt(), nextInt()-1
	for i := 0; i < N; i++ {
		distFromK = append(distFromK, -1)
	}
	dfs(K, -1, 0)

	for i := 0; i < Q; i++ {
		x, y := nextInt()-1, nextInt()-1
		fmt.Println(distFromK[x] + distFromK[y])
	}
}

type edge struct {
	to, dist int
}

func dfs(v, p, d int) {
	distFromK[v] = d
	for i := range G[v] {
		if distFromK[G[v][i].to] != -1 || G[v][i].to == p {
			continue
		}
		dfs(G[v][i].to, v, d+G[v][i].dist)
	}
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
