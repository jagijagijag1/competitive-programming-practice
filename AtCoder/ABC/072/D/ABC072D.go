package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, p := nextInt(), []int{}
	for i := 0; i < N; i++ {
		p = append(p, nextInt())
	}

	res := 0
	for i := 0; i < N-1; i++ {
		if i+1 == p[i] {
			p[i], p[i+1] = p[i+1], p[i]
			res++
		}
	}

	if p[N-1] == N {
		res++
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
