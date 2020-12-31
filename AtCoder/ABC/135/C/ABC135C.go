package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	A, B := make([]int, N+1), make([]int, N)
	for i := 0; i < N+1; i++ {
		A[i] = nextInt()
	}
	for i := 0; i < N; i++ {
		B[i] = nextInt()
	}

	res := 0
	for i := range B {
		res += min(A[i], B[i])
		B[i] -= min(A[i], B[i])
		// if A[i] <= B[i] {
		// 	res += A[i]
		// 	B[i] -= A[i]
		// } else {
		// 	res += B[i]
		// 	B[i] = 0
		// }

		res += min(A[i+1], B[i])
		A[i+1] -= min(A[i+1], B[i])
		// if B[i] <= A[i+1] {
		// 	res += B[i]
		// 	A[i+1] -= B[i]
		// } else {
		// 	res += A[i+1]
		// 	A[i+1] = 0
		// }
	}

	fmt.Println(res)
}

func min(a, b int) int {
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
