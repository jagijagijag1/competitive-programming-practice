package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

// func main() {
func mainABC054C() {
	sc.Split(bufio.ScanWords)
	N, M, a, b := nextInt(), nextInt(), []int{}, []int{}
	for i := 0; i < M; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}

	fmt.Printf("%d\n", ABC054C(N, a, b))
}

// ABC054C ...
func ABC054C(N int, a, b []int) (res int) {
	graph := []GraphNode{}
	for i := 0; i < N; i++ {
		graph = append(graph, GraphNode{i, []int{}})
	}
	for i := 0; i < len(a); i++ {
		graph[a[i]-1].edges = append(graph[a[i]-1].edges, b[i]-1)
		graph[b[i]-1].edges = append(graph[b[i]-1].edges, a[i]-1)
	}

	return dfsABC054(graph[0], graph, []int{}, 0)
}

func dfsABC054(current GraphNode, graph []GraphNode, path []int, res int) int {
	for _, n := range path {
		if n == current.id {
			return res
		}
	}

	path = append(path, current.id)
	if len(path) == len(graph) {
		return res + 1
	}

	for _, e := range current.edges {
		res = dfsABC054(graph[e], graph, path, res)
	}

	return res
}

// GraphNode ...
type GraphNode struct {
	id    int
	edges []int
}
