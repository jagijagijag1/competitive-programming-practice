package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainABC145C() {
	sc.Split(bufio.ScanWords)
	N, cties := nextInt(), []Point{}

	for i := 0; i < N; i++ {
		cties = append(cties, Point{nextInt(), nextInt()})
	}

	fmt.Printf("%f\n", ABC145C(N, cties))
}

// ABC145C ...
func ABC145C(N int, cties []Point) float64 {
	idx := []int{}
	for i := 0; i < N; i++ {
		idx = append(idx, i)
	}
	perm := permutations(idx)

	dist := float64(0.0)
	for _, p := range perm {
		for i := 0; i < len(p)-1; i++ {
			dist += pdist(cties[p[i]], cties[p[i+1]])
		}
	}

	return dist / float64(len(perm))
}

// Point ...
type Point struct {
	x, y int
}

func pdist(p1, p2 Point) float64 {
	xd := (p1.x - p2.x) * (p1.x - p2.x)
	yd := (p1.y - p2.y) * (p1.y - p2.y)
	return math.Sqrt(float64(xd + yd))
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
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
