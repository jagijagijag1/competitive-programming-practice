package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	board := [1001][1001]int{}
	for i := 0; i < n; i++ {
		lx, ly, rx, ry := nextInt()-1, nextInt()-1, nextInt()-1, nextInt()-1
		board[lx][ly]++
		board[lx][ry]--
		board[rx][ly]--
		board[rx][ry]++
	}

	for x := 0; x < 1001; x++ {
		for y := 1; y < 1001; y++ {
			board[x][y] += board[x][y-1]
		}
	}
	for x := 1; x < 1001; x++ {
		for y := 0; y < 1001; y++ {
			board[x][y] += board[x-1][y]
		}
	}

	res := make([]int, n+1)
	for x := 0; x < 1001; x++ {
		for y := 0; y < 1001; y++ {
			res[board[x][y]]++
		}
	}
	for i := 1; i < n+1; i++ {
		fmt.Println(res[i])
	}
}

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

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)

// var sc = bufio.NewScanner((os.Stdin))
