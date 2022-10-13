package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var G [][]int
var color []int
var cnt1, cnt2 int

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	G, color = make([][]int, n), make([]int, n)
	for i := range G {
		G[i] = []int{}
	}
	for i := 0; i < n-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}

	cnt1, cnt2 = 0, 0
	dfs(0, 1)

	cres := 0
	if cnt1 < cnt2 {
		cres = -1
	} else {
		cres = 1
	}
	fmt.Println(cnt1, cnt2, cres)
	fmt.Println(color)
	res := []int{}
	for i := range color {
		if color[i] == cres {
			res = append(res, i+1)
			fmt.Println(res, len(res), n/2)
			if len(res) == n/2 {
				fmt.Println("break")
				break
			}
		}
	}
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
}

func dfs(cnode, ccolor int) {
	fmt.Println(cnode, ccolor)
	if color[cnode] != 0 {
		return
	}

	color[cnode] = ccolor
	if ccolor == 1 {
		cnt1++
	} else {
		cnt2++
	}

	for _, v := range G[cnode] {
		dfs(v, ccolor*-1)
	}
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
