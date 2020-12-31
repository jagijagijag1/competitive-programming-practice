package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainAGC027A() {
	sc.Split(bufio.ScanWords)
	N, x, a := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	fmt.Printf("%d\n", AGC027A(x, a))
}

// AGC027A ...
func AGC027A(x int, a []int) (res int) {
	sort.Sort(sort.IntSlice(a))

	for i := 0; i < len(a)-1; i++ {
		if x < a[i] {
			break
		} else {
			x -= a[i]
			res++
		}
	}

	if a[len(a)-1] == x {
		res++
	}

	return
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
