package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	q := nextInt()
	m := map[int]int{}
	maxList, minList := &intDescHeap{}, &intAscHeap{}
	heap.Init(maxList)
	heap.Init(minList)

	out := bufio.NewWriter(os.Stdout)
	for i := 0; i < q; i++ {
		t := nextInt()
		if t == 1 {
			x := nextInt()
			m[x]++
			if m[x] == 1 {
				heap.Push(maxList, x)
				heap.Push(minList, x)
			}
		} else if t == 2 {
			x, c := nextInt(), nextInt()
			m[x] = max(0, m[x]-c)
			if m[x] == 0 {
			}
		} else {
			var ma, mi int
			for {
				v := heap.Pop(maxList).(int)
				if 0 < m[v] {
					ma = v
					heap.Push(maxList, ma)
					break
				}
			}
			for {
				v := heap.Pop(minList).(int)
				if 0 < m[v] {
					mi = v
					heap.Push(minList, mi)
					break
				}
			}
			fmt.Fprintf(out, "%d\n", ma-mi)
		}
	}
	defer out.Flush()
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type intAscHeap []int

func (ah intAscHeap) Len() int           { return len(ah) }
func (ah intAscHeap) Less(i, j int) bool { return ah[i] < ah[j] }
func (ah intAscHeap) Swap(i, j int)      { ah[i], ah[j] = ah[j], ah[i] }
func (ah *intAscHeap) Push(x interface{}) {
	*ah = append(*ah, x.(int))
}
func (ah *intAscHeap) Pop() interface{} {
	old := *ah
	n := len(old)
	x := old[n-1]
	*ah = old[0 : n-1]
	return x
}

type intDescHeap []int

func (dh intDescHeap) Len() int           { return len(dh) }
func (dh intDescHeap) Less(i, j int) bool { return dh[i] > dh[j] }
func (dh intDescHeap) Swap(i, j int)      { dh[i], dh[j] = dh[j], dh[i] }
func (dh *intDescHeap) Push(x interface{}) {
	*dh = append(*dh, x.(int))
}
func (dh *intDescHeap) Pop() interface{} {
	old := *dh
	n := len(old)
	x := old[n-1]
	*dh = old[0 : n-1]
	return x
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
