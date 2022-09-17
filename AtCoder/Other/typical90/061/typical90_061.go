package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	q := nextInt()
	dq := Deque()

	for i := 0; i < q; i++ {
		t, x := nextInt(), nextInt()
		if t == 1 {
			dq.PushFront(x)
		} else if t == 2 {
			dq.PushBack(x)
		} else {
			fmt.Println(dq.GetIth(x - 1))
		}
	}
}

type deque_type struct {
	Dat []int
	L   int
	R   int
}

func Deque() *deque_type {
	return &deque_type{Dat: make([]int, 200001), L: 100000, R: 100001}
}
func (dq *deque_type) PushBack(val int) {
	dq.Dat[dq.R] = val
	dq.R++
}
func (dq *deque_type) PushFront(val int) {
	dq.Dat[dq.L] = val
	dq.L--
}
func (dq *deque_type) RemoveBack() {
	dq.R--
}
func (dq *deque_type) RemoveFront() {
	dq.L++
}
func (dq *deque_type) GetBack() int {
	return dq.Dat[dq.R-1]
}
func (dq *deque_type) GetFront() int {
	return dq.Dat[dq.L+1]
}
func (dq *deque_type) GetIth(idx int) int {
	return dq.Dat[dq.L+idx+1]
}
func (dq *deque_type) Len() int {
	return dq.R - dq.L - 1
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
