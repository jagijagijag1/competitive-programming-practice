package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, A, mins, zero, sum := nextInt(), []int{}, 0, false, 0
	for i := 0; i < N; i++ {
		a := nextInt()
		if a < 0 {
			mins++
		} else if a == 0 {
			zero = true
		}

		A = append(A, abs(a))
		sum += abs(a)
	}

	if zero || mins%2 == 0 {
		fmt.Println(sum)
	} else {
		sort.Ints(A)
		fmt.Println(sum - 2*abs(A[0]))
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
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
