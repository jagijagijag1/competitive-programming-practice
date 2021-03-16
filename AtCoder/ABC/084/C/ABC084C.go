package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, C, S, F := nextInt(), []int{}, []int{}, []int{}
	for i := 0; i < N-1; i++ {
		C = append(C, nextInt())
		S = append(S, nextInt())
		F = append(F, nextInt())
	}

	for i := 0; i < N-1; i++ {
		t := S[i] + C[i]
		for j := i + 1; j < N-1; j++ {
			if t-S[j] < 0 {
				t = S[j] + C[j]
			} else if t%F[j] == 0 {
				t += C[j]
			} else {
				t += F[j] - t%F[j] + C[j]
			}
		}
		fmt.Println(t)
	}

	fmt.Println(0)
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

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
