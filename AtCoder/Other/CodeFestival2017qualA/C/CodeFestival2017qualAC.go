package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w, m := nextInt(), nextInt(), map[byte]int{}
	for i := 0; i < h; i++ {
		a := nextLine()
		for j := 0; j < w; j++ {
			m[a[j]]++
		}
	}

	c1, c2, c4 := 0, 0, 0
	for _, v := range m {
		if v%4 == 1 || v%4 == 3 {
			c1++
		} else if v%4 == 2 {
			c2++
		} else {
			c4++
		}
	}

	if h%2 == 1 && w%2 == 1 {
		c1--
	}
	if 0 < c1 {
		fmt.Println("No")
		return
	}

	if h%2 == 1 {
		c2 -= w / 2
	}
	if w%2 == 1 {
		c2 -= h / 2
	}
	if 0 < c2 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
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

// const (
// 	initialBufSize = 10000
// 	maxBufSize     = 1000000
// )

// var (
// 	sc *bufio.Scanner = func() *bufio.Scanner {
// 		sc := bufio.NewScanner(os.Stdin)
// 		buf := make([]byte, initialBufSize)
// 		sc.Buffer(buf, maxBufSize)
// 		return sc
// 	}()
// )
