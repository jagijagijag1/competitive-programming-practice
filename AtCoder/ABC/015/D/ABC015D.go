package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC015D() {
	sc.Split(bufio.ScanWords)
	W, N, K := nextInt(), nextInt(), nextInt()
	A, B := []int{}, []int{}

	for i := 0; i < N; i++ {
		A = append(A, nextInt())
		B = append(B, nextInt())
	}

	fmt.Printf("%d\n", ABC015D(W, N, K, A, B))
}

var dpABC015D [][][]int

// ABC015D ...
func ABC015D(W, N, K int, A, B []int) int {
	dpABC015D = make([][][]int, W+1)
	for i := 0; i < W+1; i++ {
		dpABC015D[i] = make([][]int, K+1)
		for j := range dpABC015D[i] {
			dpABC015D[i][j] = make([]int, N+1)
		}
	}

	return dfsABC015D(0, 0, 0, W, N, K, A, B)
}

func dfsABC015D(usedW, usedNum, current, W, N, K int, A, B []int) (res int) {
	if usedW > W || usedNum > K || current >= len(A) {
		return
	}

	if dpABC015D[usedW][usedNum][current] != 0 {
		return dpABC015D[usedW][usedNum][current]
	}

	if usedW+A[current] <= W && usedNum+1 <= K {
		res = dfsABC015D(usedW+A[current], usedNum+1, current+1, W, N, K, A, B) + B[current]
	}
	tmp := dfsABC015D(usedW, usedNum, current+1, W, N, K, A, B)

	if res < tmp {
		dpABC015D[usedW][usedNum][current] = tmp
		return tmp
	}

	dpABC015D[usedW][usedNum][current] = res
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
