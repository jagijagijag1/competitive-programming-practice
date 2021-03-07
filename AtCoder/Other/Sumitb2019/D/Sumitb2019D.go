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
	_, S, res := nextInt(), nextLine(), 0

	for i := 0; i <= 9; i++ {
		pos1 := strings.Index(S, fmt.Sprint(i))
		if pos1 == -1 {
			continue
		}

		for j := 0; j <= 9; j++ {
			subpos2 := strings.Index(S[pos1+1:], fmt.Sprint(j))
			if subpos2 == -1 {
				continue
			}

			pos2 := pos1 + subpos2 + 1
			for k := 0; k <= 9; k++ {
				subpos3 := strings.Index(S[pos2+1:], fmt.Sprint(k))
				if subpos3 == -1 {
					continue
				}

				res++
			}
		}
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
