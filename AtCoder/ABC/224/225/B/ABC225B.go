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
	t := make([][]int, n)
	for i := 0; i < n-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		t[a] = append(t[a], b)
		t[b] = append(t[b], a)
	}
	for i := range t {
		if len(t[i]) == n-1 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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

// var sc = bufio.NewScanner((os.Stdin))
