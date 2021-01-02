package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()

	if N == 1 && M == 1 {
		fmt.Println(1)
	} else if N == 1 || M == 1 {
		fmt.Println(max(N, M) - 2)
	} else {
		fmt.Println((N - 2) * (M - 2))
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
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
