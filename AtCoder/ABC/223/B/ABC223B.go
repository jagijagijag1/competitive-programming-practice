package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	s := nextLine()
	mins, maxs := s, s
	for i := 0; i < len(s); i++ {
		ss := string(s[len(s)-1]) + s[:len(s)-1]
		if strings.Compare(ss, mins) == -1 {
			mins = ss
		}
		if strings.Compare(ss, maxs) == 1 {
			maxs = ss
		}
		s = ss
	}
	fmt.Println(mins)
	fmt.Println(maxs)
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
