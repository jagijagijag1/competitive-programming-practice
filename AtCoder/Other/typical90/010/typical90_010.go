package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	cs1, cs2 := make([]int, n+1), make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if c := nextInt(); c == 1 {
			cs1[i] = nextInt()
			if i != 0 {
				cs1[i] += cs1[i-1]
				cs2[i] = cs2[i-1]
			}
		} else {
			cs2[i] = nextInt()
			if i != 0 {
				cs2[i] += cs2[i-1]
				cs1[i] = cs1[i-1]
			}
		}
	}

	q := nextInt()
	for i := 0; i < q; i++ {
		l, r := nextInt(), nextInt()
		fmt.Println(cs1[r]-cs1[l-1], cs2[r]-cs2[l-1])
	}
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
