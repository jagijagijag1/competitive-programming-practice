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
	N, a, ba, b, res := nextInt(), []string{}, []string{}, []string{}, 0
	for i := 0; i < N; i++ {
		ss := nextLine()
		res += strings.Count(ss, "AB")
		if ss[0] == 'B' && ss[len(ss)-1] == 'A' {
			ba = append(ba, ss)
		} else if ss[0] == 'B' {
			b = append(b, ss)
		} else if ss[len(ss)-1] == 'A' {
			a = append(a, ss)
		}
	}

	if len(ba) == 0 {
		res += min(len(a), len(b))
	} else if len(a)+len(b) == 0 {
		res += len(ba) - 1
	} else {
		res += len(ba) + min(len(a), len(b))
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
