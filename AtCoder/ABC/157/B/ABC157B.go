package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC157B() {
	sc.Split(bufio.ScanWords)

	A := make([][]int, 3)
	for i := range A {
		A[i] = make([]int, 3)
		for j := range A[i] {
			A[i][j] = nextInt()
		}
	}

	N, b := nextInt(), []int{}
	for i := 0; i < N; i++ {
		b = append(b, nextInt())
	}

	fmt.Printf("%s\n", ABC157B(A, b))
}

// ABC157B ...
func ABC157B(A [][]int, b []int) string {
	res := make([][]bool, 3)
	for i := range res {
		res[i] = make([]bool, 3)
	}

	for _, bi := range b {
		for i := range A {
			for j := range A[i] {
				if A[i][j] == bi {
					res[i][j] = true
				}
			}
		}
	}

	lines := [][]Point{{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}}, {{1, 0}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}},
	}
	for _, l := range lines {
		if res[l[0].y][l[0].x] && res[l[1].y][l[1].x] && res[l[2].y][l[2].x] {
			return "Yes"
		}
	}

	return "No"
}

// Point ...
type Point struct {
	x, y int
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
