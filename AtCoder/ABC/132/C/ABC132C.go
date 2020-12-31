package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC132C() {
	sc.Split(bufio.ScanWords)
	N, d := nextInt(), []int{}
	for i := 0; i < N; i++ {
		d = append(d, nextInt())
	}

	fmt.Printf("%d\n", ABC132C(d))
}

// ABC132C ...
func ABC132C(d []int) int {
	if len(d)%2 == 1 {
		return 0
	}

	sort.Sort(sort.IntSlice(d))
	m1, m2 := d[len(d)/2-1], d[len(d)/2]
	if m1 == m2 {
		return 0
	}

	res := m2 - m1

	return res
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
