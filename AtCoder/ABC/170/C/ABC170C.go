package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	x, n, p := nextInt(), nextInt(), map[int]struct{}{}
	for i := 0; i < n; i++ {
		p[nextInt()] = struct{}{}
	}

	for i := 0; i < 100; i++ {
		if _, ok := p[x-i]; !ok {
			fmt.Println(x - i)
			return
		} else if _, ok := p[x+i]; !ok {
			fmt.Println(x + i)
			return
		}
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
