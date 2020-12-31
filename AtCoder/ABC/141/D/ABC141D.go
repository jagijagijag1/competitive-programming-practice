package main

import (
	"bufio"
	"container/heap"
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
func mainABC141D() {
	sc.Split(bufio.ScanWords)
	N, M, A := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	fmt.Printf("%d\n", ABC141D(M, A))
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ABC141D ...
func ABC141D(M int, A []int) (res int) {
	ih := &intHeap{}
	heap.Init(ih)
	for i := 0; i < len(A); i++ {
		heap.Push(ih, A[i])
	}

	for i := 0; i < M; i++ {
		max := heap.Pop(ih).(int)
		heap.Push(ih, max/2)
	}

	for ih.Len() > 0 {
		res += heap.Pop(ih).(int)
	}

	return
}
