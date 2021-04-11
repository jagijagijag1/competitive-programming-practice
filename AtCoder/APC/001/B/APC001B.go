package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, a, b := nextInt(), []int{}, []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}
	for i := 0; i < N; i++ {
		b = append(b, nextInt())
	}

	total, cnt := 0, 0
	for i := 0; i < N; i++ {
		total += b[i] - a[i]
		if a[i] < b[i] {
			cnt += (b[i] - a[i] + 1) / 2
		}
	}

	if cnt <= total {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
