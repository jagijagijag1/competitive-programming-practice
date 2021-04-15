package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	s, K := nextLine(), nextInt()
	res := make([]rune, len(s))

	for i := range s {
		k := int('z' - s[i] + 1)
		if s[i] == 'a' {
			k = 0
		}

		if k <= K {
			res[i] = 'a'
			K -= k
		} else {
			res[i] = rune(s[i])
		}
	}
	if 0 < K {
		K %= 26
		res[len(res)-1] = rune(int(res[len(res)-1]) + K)
	}

	fmt.Println(string(res))
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
