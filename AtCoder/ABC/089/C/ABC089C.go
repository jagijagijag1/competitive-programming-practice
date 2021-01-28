package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, S := nextInt(), map[byte]int64{}
	for i := 0; i < N; i++ {
		s := nextLine()
		S[s[0]]++
	}

	res := int64(0)

	res += S['M'] * S['A'] * S['R']
	res += S['M'] * S['A'] * S['C']
	res += S['M'] * S['A'] * S['H']
	res += S['M'] * S['R'] * S['C']
	res += S['M'] * S['R'] * S['H']
	res += S['M'] * S['C'] * S['H']
	res += S['A'] * S['R'] * S['C']
	res += S['A'] * S['R'] * S['H']
	res += S['A'] * S['C'] * S['H']
	res += S['R'] * S['C'] * S['H']

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
