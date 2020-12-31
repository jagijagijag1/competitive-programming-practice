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
func mainABC130D() {
	sc.Split(bufio.ScanWords)
	N, K, a := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	fmt.Printf("%d\n", ABC130D(K, a))
}

// ABC130D ...
func ABC130D(K int, a []int) (res uint) {
	r, sum := 0, 0
	for l := 0; l < len(a); l++ {
		for r < len(a) && sum < K {
			sum += a[r]
			r++
		}

		if sum >= K {
			res += uint(len(a) - r + 1)
		}
		if r == l {
			r++
		} else {
			sum -= a[l]
		}
	}

	return
}

// ABC130Dnaive ...
func ABC130Dnaive(K int, a []int) (res uint) {
	dp := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		dp[i] = make([]int, len(a))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j <= i; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = a[0]
			} else {
				dp[i][j] = dp[i-1][j] + a[i]
			}
			if dp[i][j] >= K {
				res += uint(len(a) - j)
				break
			}
		}
	}

	return
}
