package main

import (
	"bufio"
	"fmt"
	"os"
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
func mainARC002D() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()
	x, y := []int{}, []int{}

	for i := 0; i < M; i++ {
		x = append(x, nextInt())
		y = append(y, nextInt())
	}

	fmt.Printf("%d\n", ABC002D(N, x, y))
}

// ABC002D ...
func ABC002D(N int, x, y []int) (res int) {
	adj := make([][]int, N)
	for i := 0; i < N; i++ {
		adj[i] = make([]int, N)
	}

	for i := 0; i < len(x); i++ {
		adj[x[i]-1][y[i]-1] = 1
		adj[y[i]-1][x[i]-1] = 1
	}

	for bit := 0; bit < (1 << uint(N)); bit++ {
		size, valid := 0, true
		for i := 0; i < N; i++ {
			if bit>>uint(i)&1 == 1 {
				size++
			}
			for j := i + 1; j < N; j++ {
				if bit>>uint(i)&1 == 1 && bit>>uint(j)&1 == 1 {
					if adj[i][j] == 0 {
						valid = false
					}
				}
			}
		}

		if valid && res < size {
			res = size
		}
	}

	return
}
