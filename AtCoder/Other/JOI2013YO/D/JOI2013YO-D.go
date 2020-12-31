package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainJOI2013YOD() {
	sc.Split(bufio.ScanWords)
	D, N, T, A, B, C := nextInt(), nextInt(), []int{}, []int{}, []int{}, []int{}
	for i := 0; i < D; i++ {
		T = append(T, nextInt())
	}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
		B = append(B, nextInt())
		C = append(C, nextInt())
	}

	fmt.Printf("%d\n", JOI2013YOD(D, N, T, A, B, C))
}

// dp[i][j]: i日目に，服jを選んだ時の，最大の派手さの絶対値
var dpJOI2013YOD [][]int

// JOI2013YOD ..j
func JOI2013YOD(D, N int, T, A, B, C []int) (res int) {
	dpJOI2013YOD = make([][]int, D)
	for i := 0; i < D; i++ {
		dpJOI2013YOD[i] = make([]int, N)
		for j := range dpJOI2013YOD[i] {
			dpJOI2013YOD[i][j] = -1
		}
	}

	// init
	for j := 0; j < N; j++ {
		if A[j] <= T[0] && T[0] <= B[j] {
			dpJOI2013YOD[0][j] = 0
		}
	}

	for i := 1; i < D; i++ {
		for j := 0; j < N; j++ {
			// day i-1, last worn j
			for k := 0; k < N; k++ {
				if dpJOI2013YOD[i-1][j] == -1 {
					continue
				}

				// day i, wear k
				if A[k] <= T[i] && T[i] <= B[k] {
					tmp := dpJOI2013YOD[i-1][j] + absInt(C[j]-C[k])

					if dpJOI2013YOD[i][k] < tmp {
						dpJOI2013YOD[i][k] = tmp
					}
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		if res < dpJOI2013YOD[D-1][i] {
			res = dpJOI2013YOD[D-1][i]
		}
	}

	return
}

func absInt(a int) int {
	if a < 0 {
		return -a
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
