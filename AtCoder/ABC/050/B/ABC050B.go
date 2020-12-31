package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, T := nextInt(), []int{}
	for i := 0; i < N; i++ {
		T = append(T, nextInt())
	}
	M, P, X := nextInt(), []int{}, []int{}
	for i := 0; i < M; i++ {
		P = append(P, nextInt())
		X = append(X, nextInt())
	}

	for i := range P {
		res := 0
		for j := range T {
			if j == P[i]-1 {
				res += X[i]
			} else {
				res += T[j]
			}
		}
		fmt.Println(res)
	}
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
