package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC141C() {
	sc.Split(bufio.ScanWords)
	N, K, Q, A := nextInt(), nextInt(), nextInt(), []int{}
	for i := 0; i < Q; i++ {
		A = append(A, nextInt())
	}

	for _, a := range ABC141C(N, K, A) {
		fmt.Printf("%s ", a)
	}
	fmt.Printf("\n")
	// fmt.Printf("%d\n", ABC141C(A))
}

// ABC141C ...
func ABC141C(N, K int, A []int) []string {
	res, ans := make([]string, N), make([]int, N)
	for _, a := range A {
		ans[a-1]++
	}

	for i := range ans {
		if K-(len(A)-ans[i]) > 0 {
			res[i] = "Yes"
		} else {
			res[i] = "No"
		}
	}

	return res
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
