package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
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
func mainABC127D() {
	sc.Split(bufio.ScanWords)
	N, M, A, CC := nextInt(), nextInt(), []int{}, []ChangeCard{}

	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}
	for i := 0; i < M; i++ {
		CC = append(CC, ChangeCard{nextInt(), nextInt()})
	}

	fmt.Printf("%d\n", ABC127D(A, CC))
}

// ABC127D ...
func ABC127D(A []int, CC []ChangeCard) (res int) {
	sort.Sort(sort.IntSlice(A))

	sort.Slice(CC, func(i, j int) bool {
		if CC[i].c > CC[j].c {
			return true
		}
		return false
	})

	D := []int{}
	for len(D) <= len(A) && len(CC) > 0 {
		cc := CC[0]
		if len(CC) != 0 {
			CC = CC[1:]
		} else {
			CC = []ChangeCard{}
		}

		for i := 0; i < cc.b && len(D) <= len(A); i++ {
			D = append(D, cc.c)
		}
	}

	for i := 0; i < len(A); i++ {
		if i >= len(D) || A[i] > D[i] {
			res += A[i]
		} else {
			res += D[i]
		}
	}

	return
}

// ABC127Dtle ...
func ABC127Dtle(A []int, CC []ChangeCard) (res int) {
	ih := &intSmallHeap{}
	heap.Init(ih)
	for i := 0; i < len(A); i++ {
		heap.Push(ih, A[i])
	}

	for i := 0; i < len(CC); i++ {
		for j := 0; j < CC[i].b; j++ {
			min := heap.Pop(ih).(int)
			if min < CC[i].c {
				heap.Push(ih, CC[i].c)
			} else {
				heap.Push(ih, min)
			}
		}
	}

	for ih.Len() > 0 {
		res += heap.Pop(ih).(int)
	}

	return
}

// ChangeCard ...
type ChangeCard struct {
	b int
	c int
}

type intSmallHeap []int

func (h intSmallHeap) Len() int           { return len(h) }
func (h intSmallHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intSmallHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intSmallHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intSmallHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
