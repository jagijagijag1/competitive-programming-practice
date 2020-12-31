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
func mainKeyence2020B() {
	sc.Split(bufio.ScanWords)
	N, X, L := nextInt(), []int64{}, []int64{}
	for i := 0; i < N; i++ {
		X = append(X, int64(nextInt()))
		L = append(L, int64(nextInt()))
	}

	fmt.Printf("%d\n", Keyence2020B(X, L))
}

// RoboArm ...
type RoboArm struct {
	pos, left, right int64
}

// Keyence2020B ...
func Keyence2020B(X, L []int64) int {
	ranges := []RoboArm{}
	for i := 0; i < len(X); i++ {
		ranges = append(ranges, RoboArm{X[i], X[i] - L[i], X[i] + L[i]})
	}
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].right < ranges[j].right {
			return true
		}
		return false
	})

	res, head := []RoboArm{}, ranges[0]
	ranges = ranges[1:]
	for len(ranges) != 0 {
		next := ranges[0]
		ranges = ranges[1:]

		if head.right <= next.left {
			res = append(res, head)
			head = next
		}
	}

	return len(res) + 1
}
