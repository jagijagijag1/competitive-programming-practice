package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M, K := nextInt(), nextInt(), nextInt()
	uf := NewUnionFind(N)
	friend := make([]int, N)
	for i := 0; i < M; i++ {
		A, B := nextInt()-1, nextInt()-1
		uf.Union(A, B)
		friend[A]++
		friend[B]++
	}
	block := make([]int, N)
	for i := 0; i < K; i++ {
		C, D := nextInt()-1, nextInt()-1
		if uf.Connected(C, D) {
			block[C]++
			block[D]++
		}
	}

	res := []int{}
	for i := 0; i < N; i++ {
		res = append(res, uf.size[uf.Root(i)]-friend[i]-block[i]-1)
	}
	fmt.Println(strings.Trim(fmt.Sprint(res), "[]"))
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

	if qRoot == pRoot {
		return
	}

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

// var sc = bufio.NewScanner((os.Stdin))

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
