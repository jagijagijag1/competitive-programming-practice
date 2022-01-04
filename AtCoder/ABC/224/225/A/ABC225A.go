package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s, m := nextLine(), map[byte]int{}
	for i := 0; i < 3; i++ {
		m[s[i]]++
	}
	for i := range m {
		if m[i] == 3 {
			fmt.Println(1)
			return
		} else if m[i] == 2 {
			fmt.Println(3)
			return
		}
	}
	fmt.Println(6)
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

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// var sc = bufio.NewScanner((os.Stdin))
