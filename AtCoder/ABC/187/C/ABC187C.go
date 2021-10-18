package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, ss := nextInt(), map[string]int{}

	for i := 0; i < n; i++ {
		s := nextLine()
		ss[s]++
		if s[0] == '!' {
			if ss[s[1:]] != 0 {
				fmt.Println(s[1:])
				return
			}
		} else {
			if ss["!"+s] != 0 {
				fmt.Println(s)
				return
			}
		}
	}
	fmt.Println("satisfiable")
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
