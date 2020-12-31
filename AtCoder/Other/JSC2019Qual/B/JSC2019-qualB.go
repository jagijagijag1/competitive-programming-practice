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
func mainJSC2019QualB() {
	sc.Split(bufio.ScanWords)
	N, K, A := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}

	fmt.Printf("%d\n", JSC2019QualB(K, A))
}

// JSC2019QualB ...
func JSC2019QualB(K int, A []int) int {
	cnt1, cnt2 := 0, 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if A[i] > A[j] {
				cnt1++
				if i < j {
					cnt2++
				}
			}
		}
	}

	mod := 1000000007
	cmb := K * (K - 1) / 2 % mod

	return (cnt2*K + cmb*cnt1) % mod
}
