package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var H, W int
var board [][]int

func main() {
	sc.Split(bufio.ScanWords)
	H, W = nextInt(), nextInt()
	board = make([][]int, H)
	for h := range board {
		board[h] = make([]int, W)
		for w := range board[h] {
			board[h][w] = nextInt() % 2
		}
	}

	res := []quadruple{}
	for h := range board {
		for w := range board[h] {
			if board[h][w] == 1 {
				if h+1 < H {
					board[h][w], board[h+1][w] = 0, 1-board[h+1][w]
					res = append(res, quadruple{h, w, h + 1, w})
				} else if w+1 < W {
					board[h][w], board[h][w+1] = 0, 1-board[h][w+1]
					res = append(res, quadruple{h, w, h, w + 1})
				}
			}
		}
	}

	fmt.Println(len(res))
	for _, r := range res {
		fmt.Println(r.y+1, r.x+1, r.yy+1, r.xx+1)
	}
}

type quadruple struct {
	y, x, yy, xx int
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
