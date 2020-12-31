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
func mainABC103D() {
	sc.Split(bufio.ScanWords)
	N, M, a, b := nextInt(), nextInt(), []int{}, []int{}
	for i := 0; i < M; i++ {
		a = append(a, nextInt())
		b = append(b, nextInt())
	}

	fmt.Printf("%d\n", ABC103D(N, a, b))
}

// ABC103D ...
func ABC103D(N int, a, b []int) int {
	ranges := []Range{}
	for i := 0; i < len(a); i++ {
		ranges = append(ranges, Range{a[i], b[i]})
	}
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].right < ranges[j].right {
			return true
		}
		return false
	})

	drop, lastDrop := []int{}, 1000000
	for len(ranges) != 0 {
		next := ranges[0]
		ranges = ranges[1:]

		if next.left < lastDrop && lastDrop <= next.right {
			continue
		} else {
			lastDrop = next.right
			drop = append(drop, lastDrop)
		}
	}

	return len(drop)
}

// Range ...
type Range struct {
	left, right int
}
