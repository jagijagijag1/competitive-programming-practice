package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s, t, res := nextLine(), nextLine(), 100000

	for i := 0; i < len(s)-len(t)+1; i++ {
		ss, cnt := s[i:i+len(t)], 0
		for j := 0; j < len(t); j++ {
			if ss[j] != t[j] {
				cnt++
			}
		}
		res = min(res, cnt)
	}
	fmt.Println(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

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
