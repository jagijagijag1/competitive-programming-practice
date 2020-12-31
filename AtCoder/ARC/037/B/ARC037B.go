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
func mainARC037B() {
	sc.Split(bufio.ScanWords)
	N, M, u, v := nextInt(), nextInt(), []int{}, []int{}
	for i := 0; i < M; i++ {
		u = append(u, nextInt())
		v = append(v, nextInt())
	}

	fmt.Printf("%d\n", ARC037B(N, u, v))
}

// ARC037B ...
func ARC037B(N int, u, v []int) (res int) {
	adjlist := make([][]int, N)
	for i := 0; i < len(u); i++ {
		adjlist[u[i]-1] = append(adjlist[u[i]-1], v[i]-1)
		adjlist[v[i]-1] = append(adjlist[v[i]-1], u[i]-1)
	}

	checked := make([]int, N)
	for i := 0; i < N; i++ {
		checked[i] = -1
	}

	for i := 0; i < N; i++ {
		if checked[i] != -1 {
			continue
		}

		var isCyclic bool
		isCyclic, checked = dfsARC037B(i, i, -1, adjlist, checked)
		if !isCyclic {
			res++
		}
	}

	return res
}

func dfsARC037B(root, current, last int, adjlist [][]int, checked []int) (bool, []int) {
	if checked[current] == root {
		return true, checked
	}

	checked[current] = root

	isCyclic := false
	for i := 0; i < len(adjlist[current]); i++ {
		if adjlist[current][i] == last {
			continue
		}

		isCyclic, checked = dfsARC037B(root, adjlist[current][i], current, adjlist, checked)
		if isCyclic {
			return true, checked
		}
	}

	return false, checked
}
