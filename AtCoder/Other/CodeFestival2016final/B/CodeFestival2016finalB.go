package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, sum, ub, res := nextInt(), 0, 0, []int{}

	for i := 1; i <= N; i++ {
		sum += i
		if N <= sum {
			ub = i
			break
		}
	}

	sum = 0
	for i := ub; 0 < i; i-- {
		if sum+i == N {
			res = append(res, i)
			for _, r := range res {
				fmt.Println(r)
			}
			return
		} else if sum+i < N {
			res = append(res, i)
			sum += i
		}
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
