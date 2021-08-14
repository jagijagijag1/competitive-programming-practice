package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	c := []int{nextInt(), nextInt(), nextInt(), nextInt()}

	for bit := 0; bit < (1 << uint(4)); bit++ {
		s1, s2 := 0, 0
		for i := 0; i < 4; i++ {
			if bit>>uint(i)&1 == 1 {
				s1 += c[i]
			} else {
				s2 += c[i]
			}
		}
		if s1 == s2 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
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
