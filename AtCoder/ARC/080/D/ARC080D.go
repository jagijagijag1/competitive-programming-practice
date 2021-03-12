package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, W, N, a := nextInt(), nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	board := make([][]int, H)
	for i := range board {
		board[i] = make([]int, W)
	}

	color := 0
	for i := range board {
		if i%2 == 0 {
			for j := 0; j < W; j++ {
				if a[color] == 0 {
					color++
				}
				board[i][j] = color + 1
				a[color]--
			}
		} else {
			for j := W - 1; 0 <= j; j-- {
				if a[color] == 0 {
					color++
				}
				board[i][j] = color + 1
				a[color]--
			}
		}
	}

	for i := range board {
		fmt.Println(strings.Trim(fmt.Sprint(board[i]), "[]"))
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
