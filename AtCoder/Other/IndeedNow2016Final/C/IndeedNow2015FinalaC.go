package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainIndeedNow2015FinalaC() {
	sc.Split(bufio.ScanWords)
	N, M, companies, applicants := nextInt(), nextInt(), []IndeedNowCompany{}, []IndeedNowApplicant{}
	for i := 0; i < N; i++ {
		c := IndeedNowCompany{nextInt(), nextInt(), nextInt(), nextInt()}
		companies = append(companies, c)
	}
	for i := 0; i < M; i++ {
		a := IndeedNowApplicant{nextInt(), nextInt(), nextInt()}
		applicants = append(applicants, a)
	}

	// fmt.Printf("%d\n", IndeedNow2015FinalaC(N, M, companies, applicants))
	for _, r := range IndeedNow2015FinalaC(N, M, companies, applicants) {
		fmt.Printf("%d\n", r)
	}
}

// IndeedNowCompany ...
type IndeedNowCompany struct {
	a, b, c, w int
}

// IndeedNowApplicant ...
type IndeedNowApplicant struct {
	x, y, z int
}

// IndeedNow2015FinalaC ...
func IndeedNow2015FinalaC(N, M int, companies []IndeedNowCompany, applicants []IndeedNowApplicant) []int {
	dp := [101][101][101]int{}

	sort.Slice(companies, func(i, j int) bool {
		if companies[i].w < companies[j].w {
			return true
		}
		return false
	})

	for _, c := range companies {
		dp[c.a][c.b][c.c] = c.w
	}

	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				w1, w2, w3 := dp[maxInt(0, i-1)][j][k], dp[i][maxInt(0, j-1)][k], dp[i][j][maxInt(0, k-1)]
				dp[i][j][k] = maxIntInSlice([]int{dp[i][j][k], w1, w2, w3})
			}
		}
	}

	res := []int{}
	for _, a := range applicants {
		res = append(res, dp[a.x][a.y][a.z])
	}

	return res
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxIntInSlice(v []int) (res int) {
	res = math.MaxInt32
	for i, e := range v {
		if i == 0 || e < res {
			res = e
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
