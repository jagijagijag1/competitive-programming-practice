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
	b, res := make([]int, N), []int{}
	for i := 0; i < N; i++ {
		b[i] = nextInt()
	}

	for 0 < len(b) {
		isRemoved := false
		for i := len(b) - 1; 0 <= i; i-- {
			if b[i] == i+1 {
				res = append(res, b[i])
				b = append(b[0:i], b[i+1:]...)
				isRemoved = true
				break
			}
		}
		if len(b) != 0 && !isRemoved {
			fmt.Println(-1)
			return
		}
	}
	for i := N - 1; 0 <= i; i-- {
		fmt.Println(res[i])
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
