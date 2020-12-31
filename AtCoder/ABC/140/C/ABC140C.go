package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, B := nextInt(), []int{}
	for i := 0; i < N-1; i++ {
		B = append(B, nextInt())
	}

	A := make([]int, len(B)+1)
	for i := range A {
		if i == 0 {
			A[i] = B[i]
		} else if i == len(A)-1 {
			A[i] = B[i-1]
		} else {
			A[i] = minInt(B[i-1], B[i])
		}
	}

	res := 0
	for _, a := range A {
		res += a
	}
	fmt.Println(res)
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
