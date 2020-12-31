package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC139C() {
	// sc.Split(bufio.ScanWords)
	// N, v := nextInt(), []int{}
	var N int
	fmt.Scanf("%d", &N)

	H := make([]int64, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &H[i])
	}

	fmt.Printf("%d\n", ABC139C(H))
}

// ABC139C ...
func ABC139C(H []int64) (res int) {
	last, curr := H[0], 0
	for i := 1; i < len(H); i++ {
		if H[i] <= last {
			curr++
		} else {
			res = maxInt(res, curr)
			curr = 0
		}
		last = H[i]
	}

	return maxInt(res, curr)
}

func maxInt(a, b int) int {
	if a < b {
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
