package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()

	e, re := make([][]int, n), make([][]int, n)
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		e[a] = append(e[a], b)
		re[b] = append(re[b], a)
	}

	// scc step1
	order := []int{}
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		order = dfs(i, e, visited, order)
	}

	// scc step2
	cnt := []int{}
	visited = make([]bool, n)
	for i := n - 1; 0 <= i; i-- {
		o := order[i]
		if visited[o] {
			continue
		}

		p := dfs(o, re, visited, []int{})
		cnt = append(cnt, len(p))
	}

	res := 0
	for i := range cnt {
		res += cnt[i] * (cnt[i] - 1) / 2
	}
	fmt.Println(res)
}

func dfs(x int, e [][]int, visited []bool, order []int) []int {
	if visited[x] {
		return order
	}
	visited[x] = true

	for i := range e[x] {
		nx := e[x][i]
		if visited[nx] {
			continue
		}
		order = dfs(nx, e, visited, order)
	}
	order = append(order, x)
	return order
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
