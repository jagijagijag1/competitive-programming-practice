package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w, q := nextInt(), nextInt(), nextInt()
	board := make([]int, h*w) // 0: white, 1:red
	uf := NewUnionFind(h * w)
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	for i := 0; i < q; i++ {
		if nextInt() == 1 {
			r, c := nextInt()-1, nextInt()-1
			board[r*w+c] = 1
			for d := 0; d < 4; d++ {
				nr, nc := r+dx[d], c+dy[d]
				if nr < 0 || h <= nr {
					continue
				}
				if nc < 0 || w <= nc {
					continue
				}
				if board[nr*w+nc] == 1 {
					uf.Union(r*w+c, nr*w+nc)
				}
			}
		} else {
			ra, ca, rb, cb := nextInt()-1, nextInt()-1, nextInt()-1, nextInt()-1

			if board[ra*w+ca] == 0 || board[rb*w+cb] == 0 {
				fmt.Println("No")
				continue
			}

			if uf.Connected(ra*w+ca, rb*w+cb) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
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
