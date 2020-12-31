package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainARC153E() {
	sc.Split(bufio.ScanWords)
	H, N := nextInt(), nextInt()
	A, B := []int{}, []int{}

	for i := 0; i < N; i++ {
		A = append(A, nextInt())
		B = append(B, nextInt())
	}

	fmt.Printf("%d\n", ABC153E(H, N, A, B))
}

var dpABC153E [][][]int

// ABC153E ...
func ABC153E(H, N int, A, B []int) int {
	INF := math.MaxInt32

	// dp[i][j]: i番目までの魔法を任意の回数使った時に，ダメージの和をj以上にする最小コスト
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, H+1)
	}
	for i := range dp {
		for j := range dp[i] {
			if j == 0 {
				continue
			}
			dp[i][j] = INF
		}
	}

	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j <= H; j++ {
			lastj := maxInt(j-A[i], 0)
			dp[i+1][j] = minInt(dp[i][j], dp[i+1][lastj]+B[i])
		}
	}

	return dp[N][H]
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
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
