package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC147C() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	x, y := make([][]int, N), make([][]int, N)

	for i := 0; i < N; i++ {
		Ai := nextInt()
		x[i], y[i] = make([]int, Ai), make([]int, Ai)
		for j := 0; j < Ai; j++ {
			x[i][j] = nextInt()
			y[i][j] = nextInt()
		}
	}

	fmt.Printf("%d\n", ABC147C(N, x, y))
}

// ABC147C ...
func ABC147C(N int, x, y [][]int) (res int) {
	// 0: liar, 1: honest
	for bit := 0; bit < (1 << N); bit++ {
		isConsistent, tmpcnt := true, 0
		for i := 0; i < N; i++ {
			if bit>>uint(i)&1 == 0 {
				continue
			}
			tmpcnt++

			if !isConsistent {
				continue
			}

			for j := range x[i] {
				if bit>>uint(x[i][j]-1)&1 != y[i][j] {
					isConsistent = false
					break
				}
			}
		}

		if isConsistent && res < tmpcnt {
			res = tmpcnt
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
