package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	W, H, N := nextInt(), nextInt(), nextInt()
	isblack := make([][]bool, W)
	for j := range isblack {
		isblack[j] = make([]bool, H)
	}

	for i := 0; i < N; i++ {
		x, y, a := nextInt(), nextInt(), nextInt()
		switch a {
		case 1:
			for i := 0; i < x; i++ {
				for j := 0; j < H; j++ {
					isblack[i][j] = true
				}
			}
		case 2:
			for i := x; i < W; i++ {
				for j := 0; j < H; j++ {
					isblack[i][j] = true
				}
			}
		case 3:
			for i := 0; i < W; i++ {
				for j := 0; j < y; j++ {
					isblack[i][j] = true
				}
			}
		case 4:
			for i := 0; i < W; i++ {
				for j := y; j < H; j++ {
					isblack[i][j] = true
				}
			}
		}
	}

	res := 0
	for i := range isblack {
		for j := range isblack[i] {
			if !isblack[i][j] {
				res++
			}
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
