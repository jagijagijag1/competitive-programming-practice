package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	S, T := nextLine(), nextLine()
	s, t := make([]int, 26), make([]int, 26)
	for i := 0; i < 26; i++ {
		s[i], t[i] = -1, -1
	}

	for i := range S {
		si, ti := int(S[i]-'a'), int(T[i]-'a')

		if (s[si] != -1 || t[ti] != -1) && (s[si] != ti || t[ti] != si) {
			fmt.Println("No")
			return
		}

		s[si], t[ti] = ti, si
	}

	fmt.Println("Yes")
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
