package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	s := nextLine()

	occ := calcNextCharIndex(s)

	res := []rune{}
	cur := 0
	for i := 0; i < k; i++ {
		for j := 0; j < 26; j++ {
			next := occ[cur][j]
			if k-i <= n-next {
				res = append(res, rune('a'+j))
				cur = next + 1
				break
			}
		}
	}
	fmt.Println(string(res))

}

func calcNextCharIndex(s string) [][26]int {
	n := len(s)
	res := make([][26]int, n+1)
	for i := 0; i < 26; i++ {
		res[n][i] = n
	}

	for i := n - 1; 0 <= i; i-- {
		for j := 0; j < 26; j++ {
			if int(s[i]-'a') == j {
				res[i][j] = i
			} else {
				res[i][j] = res[i+1][j]
			}
		}
	}

	return res
}

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

// var sc = bufio.NewScanner((os.Stdin))
