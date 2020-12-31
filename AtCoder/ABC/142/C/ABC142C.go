package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC142C() {
	sc.Split(bufio.ScanWords)
	N, A := nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	for _, a := range ABC142C(A) {
		fmt.Printf("%d ", a)
	}
	fmt.Printf("\n")
	// fmt.Printf("%d\n", ABC142C(A))
}

// ABC142C ...
func ABC142C(A []int) (res []int) {
	// key: order, val: id
	attend := map[int]int{}
	for i := range A {
		attend[A[i]] = i + 1
	}

	for i := 1; i <= len(A); i++ {
		res = append(res, attend[i])
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
