package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, last := nextInt(), nextInt()
	for i := 1; i < N; i++ {
		h := nextInt()
		if h < last && h-1 < last {
			fmt.Println("No")
			return
		}
		last = maxInt(last, h-1)
	}
	fmt.Println("Yes")
}

func maxInt(a, b int) int {
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
