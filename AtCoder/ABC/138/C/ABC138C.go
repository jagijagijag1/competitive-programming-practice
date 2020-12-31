package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC138C() {
	sc.Split(bufio.ScanWords)
	N, v := nextInt(), []int{}
	for i := 0; i < N; i++ {
		v = append(v, nextInt())
	}

	fmt.Printf("%f\n", ABC138C(v))
}

// ABC138C ...
func ABC138C(v []int) float64 {
	fh := &float64Heap{}
	heap.Init(fh)

	for _, vv := range v {
		heap.Push(fh, float64(vv))
	}

	for fh.Len() > 1 {
		vv1 := heap.Pop(fh).(float64)
		vv2 := heap.Pop(fh).(float64)
		tmp := (vv1 + vv2) / 2.0

		heap.Push(fh, tmp)
	}

	return heap.Pop(fh).(float64)
}

type float64Heap []float64

func (h float64Heap) Len() int           { return len(h) }
func (h float64Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h float64Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *float64Heap) Push(x interface{}) {
	*h = append(*h, x.(float64))
}
func (h *float64Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
