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
	N, K, D := nextInt(), nextInt(), [10]bool{}
	for i := 0; i < K; i++ {
		D[nextInt()] = true
	}

	for i := N; i <= 10*N; i++ {
		contains := false
		for j := 0; j < 10; j++ {
			si := strconv.Itoa(i)
			if D[j] && strings.Contains(si, strconv.Itoa(j)) {
				contains = true
				break
			}
		}
		if !contains {
			fmt.Println(i)
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
