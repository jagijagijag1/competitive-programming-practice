package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainTDPCD() {
	sc.Split(bufio.ScanWords)
	N, D := nextInt(), nextInt()

	fmt.Printf("%f\n", TDPCD(N, D))
}

// TDPCD ...
func TDPCD(N, D int) float64 {
	n2, n3, n5 := 0, 0, 0
	for D%2 == 0 {
		D /= 2
		n2++
	}
	for D%3 == 0 {
		D /= 3
		n3++
	}
	for D%5 == 0 {
		D /= 5
		n5++
	}
	if D != 1 {
		return 0.0
	}

	var dp [100][100][50][50]float64
	dp[0][0][0][0] = 1

	for i := 0; i < N; i++ {
		for a := 0; a <= n2; a++ {
			for b := 0; b <= n3; b++ {
				for c := 0; c <= n5; c++ {
					dp[i+1][a][b][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][minInt(a+1, n2)][b][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][a][minInt(b+1, n3)][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][minInt(a+2, n2)][b][c] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][a][b][minInt(c+1, n5)] += (1.0 / 6) * dp[i][a][b][c]
					dp[i+1][minInt(a+1, n2)][minInt(b+1, n3)][c] += (1.0 / 6) * dp[i][a][b][c]
				}
			}
		}
	}

	return dp[N][n2][n3][n5]
}

// TDPCDwa ...
func TDPCDwa(N, D int) float64 {
	dicefactor := []int{2, 3, 5}
	factorcnt := []int{}

	for _, df := range dicefactor {
		cnt := 0
		for D%df == 0 {
			D /= df
			cnt++
		}
		factorcnt = append(factorcnt, cnt)
	}
	if D != 1 {
		return 0
	}
	n2, n3, n5 := factorcnt[0], factorcnt[1], factorcnt[2]

	dp := make([][][][]float64, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][][]float64, n2+1)
		for j2 := range dp[i] {
			dp[i][j2] = make([][]float64, n3+1)
			for j3 := range dp[i][j2] {
				dp[i][j2][j3] = make([]float64, n5+1)
			}
		}
	}

	dp[0][0][0][0] = 1
	for i := 0; i < N; i++ {
		for j2 := 0; j2 <= n2; j2++ {
			for j3 := 0; j3 <= n3; j3++ {
				for j5 := 0; j5 <= n5; j5++ {
					dp[i+1][j2][j3][j5] += dp[i][j2][j3][j5] / 6
					dp[i+1][minInt(j2+1, n2)][j3][j5] += dp[i][j2][j3][j5] / 6
					dp[i+1][j2][minInt(j3+1, n3)][j5] += dp[i][j2][j3][j5] / 6
					dp[i+1][minInt(j2+2, n2)][j3][j5] += dp[i][j2][j3][j5] / 6
					dp[i+1][j2][j3][minInt(j5+1, n5)] += dp[i][j2][j3][j5] / 6
					dp[i+1][minInt(j2+1, n2)][minInt(j3+1, n3)][j5] += dp[i][j2][j3][j5] / 6
				}
			}
		}
	}

	return dp[N][n2][n3][n5]
}

func minInt(a, b int) int {
	if a > b {
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
