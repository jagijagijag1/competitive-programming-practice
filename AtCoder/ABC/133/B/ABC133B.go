package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainABC133B() {
	sc.Split(bufio.ScanWords)
	N, D := nextInt(), nextInt()
	X := make([][]int, N)
	for i := 0; i < N; i++ {
		X[i] = make([]int, D)
		for j := 0; j < D; j++ {
			X[i][j] = nextInt()
		}
	}

	fmt.Printf("%d\n", ABC133B(X))
}

// ABC133B ...
func ABC133B(X [][]int) (res int) {
	for i := 0; i < len(X); i++ {
		for j := i + 1; j < len(X); j++ {
			sum := 0
			for k := 0; k < len(X[i]); k++ {
				sum += (X[i][k] - X[j][k]) * (X[i][k] - X[j][k])
			}
			dist := math.Sqrt(float64(sum))
			if dist == float64(int(dist)) {
				res++
			}
		}
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
