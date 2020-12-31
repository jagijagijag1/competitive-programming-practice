package main

import (
	"bufio"
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
func mainARC029A() {
	sc.Split(bufio.ScanWords)
	N, t := nextInt(), []int{}
	for i := 0; i < N; i++ {
		t = append(t, nextInt())
	}

	fmt.Printf("%d\n", ARC029A(t))
}

// ARC029A ...
func ARC029A(t []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(t)))

	t1, t2 := 0, 0
	for i := 0; i < len(t); i++ {
		if t1 < t2 {
			t1 += t[i]
		} else {
			t2 += t[i]
		}
	}

	if t1 < t2 {
		return t2
	}
	return t1

}
