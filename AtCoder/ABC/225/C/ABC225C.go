package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()

	last := make([]int, m)
	for i := 0; i < n; i++ {
		current := make([]int, m)
		for j := 0; j < m; j++ {
			current[j] = nextInt()
			if j != 0 && current[j]-current[j-1] != 1 {
				fmt.Println("No")
				return
			}
			if j != m-1 && m != 1 && current[j]%7 == 0 {
				fmt.Println("No")
				return
			}
		}

		for j := 0; j < m; j++ {
			if i != 0 && current[j]-last[j] != 7 {
				fmt.Println("No")
				return
			}
			last[j] = current[j]
		}
	}
	fmt.Println("Yes")
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
