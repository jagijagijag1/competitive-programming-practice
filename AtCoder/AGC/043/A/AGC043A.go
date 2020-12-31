package main

import (
	"bufio"
	"fmt"
	"math"
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
func mainAGC043A() {
	sc.Split(bufio.ScanWords)
	H, W := nextInt(), nextInt()
	S := []string{}
	for i := 0; i < H; i++ {
		S = append(S, nextLine())
	}

	fmt.Printf("%d\n", AGC043A(H, W, S))
}

// AGC043A ...
func AGC043A(H, W int, S []string) int {
	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if i == 0 && j == 0 {
				if S[0][0] == '#' {
					dp[0][0] = 1
				} else {
					dp[0][0] = 0
				}
			} else {
				up, left := math.MaxInt32, math.MaxInt32
				if i != 0 {
					up = dp[i-1][j]
					if S[i-1][j] == '.' && S[i][j] == '#' {
						up++
					}
				}
				if j != 0 {
					left = dp[i][j-1]
					if S[i][j-1] == '.' && S[i][j] == '#' {
						left++
					}
				}

				if up < left {
					dp[i][j] = up
				} else {
					dp[i][j] = left
				}
			}
		}
	}

	return dp[H-1][W-1]
}
