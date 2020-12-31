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
func mainABC106D() {
	sc.Split(bufio.ScanWords)
	N, M, Q := nextInt(), nextInt(), nextInt()
	L, R, p, q := []int{}, []int{}, []int{}, []int{}

	for i := 0; i < M; i++ {
		L = append(L, nextInt())
		R = append(R, nextInt())
	}

	for i := 0; i < Q; i++ {
		p = append(p, nextInt())
		q = append(q, nextInt())
	}

	//fmt.Printf("%d\n", ABC106D(N, L, R, p, q))
	res := ABC106D(N, L, R, p, q)
	for _, r := range res {
		fmt.Println(r)
	}
}

// ABC106D ...
func ABC106D(N int, L, R, p, q []int) []int {
	res := []int{}

	// init
	sumtable := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		sumtable[i] = make([]int, N+1)
	}
	for i := 0; i < len(L); i++ {
		sumtable[L[i]][R[i]]++
	}

	// cum sum
	for i := 1; i < N+1; i++ {
		for j := 1; j < N+1; j++ {
			sumtable[i][j] += sumtable[i-1][j]
			sumtable[i][j] += sumtable[i][j-1]
			sumtable[i][j] -= sumtable[i-1][j-1]
		}
	}

	// query
	for i := 0; i < len(p); i++ {
		tmp := sumtable[q[i]][q[i]] - sumtable[q[i]][p[i]-1] - sumtable[p[i]-1][q[i]] + sumtable[p[i]-1][p[i]-1]
		res = append(res, tmp)
	}

	return res
}
