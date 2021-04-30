package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	H, W, A, B := nextInt(), nextInt(), nextInt(), nextInt()
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if h < B {
				if w < A {
					fmt.Print(0)
				} else {
					fmt.Print(1)
				}
			} else {
				if w < A {
					fmt.Print(1)
				} else {
					fmt.Print(0)
				}
			}
		}
		fmt.Println()
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
