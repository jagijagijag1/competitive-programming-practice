package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s, t := []rune(nextLine()), nextLine()
	if string(s) == t {
		fmt.Println("Yes")
		return
	}

	for i := 0; i < len(s)-1; i++ {
		s[i], s[i+1] = s[i+1], s[i]
		if string(s) == t {
			fmt.Println("Yes")
			return
		}
		s[i], s[i+1] = s[i+1], s[i]
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

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// var sc = bufio.NewScanner((os.Stdin))
