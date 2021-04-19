package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, m := nextInt(), map[int]int{}
	for i := 0; i < N; i++ {
		m[nextInt()]++
	}

	res, cnt2 := 0, 0
	for _, v := range m {
		if v%2 == 1 {
			res++
		} else {
			cnt2++
		}
	}

	res += cnt2
	if cnt2%2 == 1 {
		res--
	}
	fmt.Println(res)
}

// var sc = bufio.NewScanner((os.Stdin))

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

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)
