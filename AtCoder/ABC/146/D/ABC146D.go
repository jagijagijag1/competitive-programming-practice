package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var E [][]int
var eid [][]int
var color []int

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	E, eid = make([][]int, N), make([][]int, N)
	for i := range E {
		E[i] = []int{}
		eid[i] = []int{}
	}
	for i := 0; i < N-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		E[a], eid[a] = append(E[a], b), append(eid[a], i)
		E[b], eid[b] = append(E[b], a), append(eid[b], i)
	}

	maxc := 0
	for i := range E {
		maxc = max(maxc, len(E[i]))
	}
	fmt.Println(maxc)

	color = make([]int, N-1)
	dfs(0, -1, -1)
	for _, c := range color {
		fmt.Println(c)
	}
}

func dfs(v, c, p int) {
	k := 1
	for i := range E[v] {
		if E[v][i] == p {
			continue
		}
		if c == k {
			k++
		}
		color[eid[v][i]] = k
		k++
		dfs(E[v][i], color[eid[v][i]], v)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
