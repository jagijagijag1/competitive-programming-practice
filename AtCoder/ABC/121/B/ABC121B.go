package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC121B() {
	sc.Split(bufio.ScanWords)
	N, M, C := nextInt(), nextInt(), nextInt()
	A, B := make([][]int, N), make([]int, M)

	for i := range B {
		B[i] = nextInt()
	}
	for i := range A {
		A[i] = make([]int, M)
		for j := range A[i] {
			A[i][j] = nextInt()
		}
	}

	fmt.Printf("%d\n", ABC121B(A, B, C))
}

// ABC121B ...
func ABC121B(A [][]int, B []int, C int) (res int) {
	for i := range A {
		sum := 0
		for j := range A[i] {
			sum += A[i][j] * B[j]
		}

		if sum+C > 0 {
			res++
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
