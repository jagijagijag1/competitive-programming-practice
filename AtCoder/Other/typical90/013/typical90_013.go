package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	G := NewGraph(n)

	for i := 0; i < m; i++ {
		a, b, c := nextInt()-1, nextInt()-1, nextInt()
		G.AddWeightedEdge(a, b, c)
		G.AddWeightedEdge(b, a, c)
	}

	distFrom1 := G.dijkstra(0)
	distFromN := G.dijkstra(n - 1)
	for i := 0; i < n; i++ {
		fmt.Println(distFrom1[i] + distFromN[i])
	}
}

type Graph struct {
	edges [][]Edge
}

type Edge struct {
	to, weight int
}

func NewGraph(n int) *Graph {
	return &(Graph{make([][]Edge, n)})
}

func (g *Graph) AddEdge(s, t int) {
	g.AddWeightedEdge(s, t, 1)
}

func (g *Graph) AddWeightedEdge(s, t, w int) {
	g.edges[s] = append(g.edges[s], Edge{t, w})
}

type DijkstraNode struct {
	node, cost int
}
type DijkstraPriorityQueue []*DijkstraNode

func (pq DijkstraPriorityQueue) Len() int            { return len(pq) }
func (pq DijkstraPriorityQueue) Less(i, j int) bool  { return pq[i].cost < pq[j].cost }
func (pq DijkstraPriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *DijkstraPriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*DijkstraNode)) }
func (pq *DijkstraPriorityQueue) Pop() interface{} {
	o := *pq
	n := len(o) - 1
	item := o[n]
	*pq = o[0:n]
	return item
}

func (g *Graph) dijkstra(s int) []int {
	n := len(g.edges)
	pq := make(DijkstraPriorityQueue, 0)
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		var c int
		if i == s {
			c = 0
		} else {
			c = int(1e18)
		}
		cost[i] = c
		heap.Push(&pq, &DijkstraNode{i, c})
	}

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(*DijkstraNode)
		for i := 0; i < len(g.edges[u.node]); i++ {
			v := g.edges[u.node][i]
			c := cost[u.node] + v.weight
			if cost[v.to] > c {
				cost[v.to] = c
				heap.Push(&pq, &DijkstraNode{v.to, c})
			}
		}
	}

	return cost
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
