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
	_, A, B, C, D, S := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextLine()

	if !strings.Contains(S[A-1:C], "##") && !strings.Contains(S[B-1:D], "##") {
		if C < D {
			fmt.Println("Yes")
			return
		} else if strings.Contains(S[B-2:D+1], "...") {
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
