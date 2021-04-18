package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M, res := nextInt(), nextInt(), 0
	a, b := make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		a[i] = nextInt() - 1
		b[i] = nextInt() - 1
	}

	for i := 0; i < M; i++ {
		uf := NewUnionFind(N)
		for j := 0; j < M; j++ {
			if i == j {
				continue
			}
			uf.Union(a[j], b[j])
		}

		r := 0
		for k := 0; k < N; k++ {
			if uf.Root(k) == k {
				r++
			}
		}
		if 1 < r {
			res++
		}
	}

	fmt.Println(res)
}

type UnionFind struct {
	root []int
	size []int
}

func NewUnionFind(size int) *UnionFind {
	return new(UnionFind).init(size)
}

func (uf *UnionFind) init(size int) *UnionFind {
	uf = new(UnionFind)
	uf.root = make([]int, size)
	uf.size = make([]int, size)

	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Union(p int, q int) {
	qRoot := uf.Root(q)
	pRoot := uf.Root(p)

	if uf.size[qRoot] < uf.size[pRoot] {
		uf.root[qRoot] = uf.root[pRoot]
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.root[pRoot] = uf.root[qRoot]
		uf.size[qRoot] += uf.size[pRoot]
	}
}

func (uf *UnionFind) Root(p int) int {
	if p > len(uf.root)-1 {
		return -1
	}

	for uf.root[p] != p {
		uf.root[p] = uf.root[uf.root[p]]
		p = uf.root[p]
	}

	return p
}

func (uf *UnionFind) Find(p int) int {
	return uf.Root(p)
}

func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
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
