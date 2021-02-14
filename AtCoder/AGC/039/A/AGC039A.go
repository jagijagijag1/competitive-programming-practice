package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	S, K := nextLine(), nextInt()

	last, rep := S[0], true
	for i := 1; i < len(S); i++ {
		if last != S[i] {
			rep = false
			break
		}
	}

	if rep {
		fmt.Println(len(S) * K / 2)
	} else if S[0] != S[len(S)-1] {
		fmt.Println(replaceCount(S) * K)
	} else {
		sum1, sum2 := replaceCount(S), replaceCount(S+S)
		fmt.Println(sum1 + (K-1)*(sum2-sum1))
	}
}

func replaceCount(S string) int {
	res, last, replaced := 0, S[0], string(S[0])
	for i := 1; i < len(S); i++ {
		if last == S[i] {
			res++
			replaced += "0"
			last = '0'
		} else {
			replaced += string(S[i])
			last = S[i]
		}
	}
	return res
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
