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

	if m[0] == N {
		fmt.Println("Yes")
		return
	}

	if N%3 != 0 {
		fmt.Println("No")
		return
	}

	if len(m) == 2 {
		for k, v := range m {
			if k == 0 && v != N/3 {
				fmt.Println("No")
				return
			}
			if k != 0 && v != 2*N/3 {
				fmt.Println("No")
				return
			}
		}
		fmt.Println("Yes")
		return
	}

	if len(m) == 3 {
		s := 0
		for k, v := range m {
			if v != N/3 {
				fmt.Println("No")
				return
			}
			s ^= k
		}

		if s == 0 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
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
