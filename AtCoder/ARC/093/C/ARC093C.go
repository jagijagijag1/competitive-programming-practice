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
	A, diff, sum := make([]int, N+2), make([]int, N+1), 0
	A[0], A[N] = 0, 0
	for i := 1; i < N+1; i++ {
		A[i] = nextInt()
		diff[i-1] = A[i] - A[i-1]
		sum += abs(diff[i-1])
	}
	diff[N] = A[N+1] - A[N]
	sum += abs(diff[N])

	for i := 0; i < N; i++ {
		if diff[i]*diff[i+1] < 0 {
			fmt.Println(sum - abs(diff[i+1]) - abs(diff[i]) + abs(diff[i+1]+diff[i]))
		} else {
			fmt.Println(sum)
		}
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

// var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

// func readLine() string {
// 	buf := make([]byte, 0, 1000000)
// 	for {
// 		l, p, e := rdr.ReadLine()
// 		if e != nil {
// 			panic(e)
// 		}
// 		buf = append(buf, l...)
// 		if !p {
// 			break
// 		}
// 	}
// 	return string(buf)
// }
