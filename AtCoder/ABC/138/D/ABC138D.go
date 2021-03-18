package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, Q := nextInt(), nextInt()
	nodes, visited = make([]node, N+1), make([]bool, N+1)
	for i := range nodes {
		nodes[i] = node{i, 0, []int{}}
	}
	for i := 0; i < N-1; i++ {
		a, b := nextInt(), nextInt()
		nodes[a].edges = append(nodes[a].edges, b)
		nodes[b].edges = append(nodes[b].edges, a)
	}

	for i := 0; i < Q; i++ {
		p, x := nextInt(), nextInt()
		nodes[p].count = nodes[p].count + x
	}

	dfs(1, 0)

	for i := 1; i < N; i++ {
		fmt.Print(nodes[i].count, " ")
	}
	fmt.Println(nodes[N].count)
}

var nodes []node
var visited []bool

type node struct {
	id    int
	count int
	edges []int
}

func dfs(n, pc int) {
	if visited[n] {
		return
	}

	nodes[n].count = nodes[n].count + pc
	visited[n] = true
	if len(nodes[n].edges) == 0 {
		return
	}

	for _, c := range nodes[n].edges {
		dfs(c, nodes[n].count)
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
