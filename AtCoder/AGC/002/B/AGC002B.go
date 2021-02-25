package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()
	num, red := make([]int, N), make([]bool, N)
	red[0] = true
	for i := 0; i < N; i++ {
		num[i] = 1
	}

	for i := 0; i < M; i++ {
		x, y := nextInt()-1, nextInt()-1
		if red[x] {
			if num[x] == 1 {
				red[x] = false
			}
			red[y] = true
		}
		num[x]--
		num[y]++
	}

	res := 0
	for i := 0; i < N; i++ {
		if red[i] {
			res++
		}
	}
	fmt.Println(res)
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
