package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainTDPCA() {
	sc.Split(bufio.ScanWords)
	N, p := nextInt(), []int{}
	for i := 0; i < N; i++ {
		p = append(p, nextInt())
	}

	fmt.Printf("%d\n", TDPCA(p))
}

// TDPCA ...
func TDPCA(p []int) (res int) {
	dp := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		dp[i] = make([]bool, 10001)
	}

	dp[0][0] = true
	for i := 1; i < len(p)+1; i++ {
		for j := 0; j < len(dp[i]); j++ {
			if dp[i-1][j] == true {
				dp[i][j] = true
				dp[i][j+p[i-1]] = true
			}
		}
	}

	for _, s := range dp[len(p)] {
		if s == true {
			res++
		}
	}

	return
}

// TDPCAwa ...
func TDPCAwa(p []int) int {
	res := map[int]struct{}{}
	for bit := 0; bit < (1 << uint(len(p))); bit++ {
		score := 0
		for i := 0; i < len(p); i++ {
			if bit>>uint(i)&1 == 1 {
				score += p[i]
			}
		}
		res[score] = struct{}{}
	}

	return len(res)
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
