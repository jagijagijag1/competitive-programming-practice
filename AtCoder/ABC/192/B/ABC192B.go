package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s := nextLine()
	for i := 0; i < len(s); i++ {
		if i%2 == 0 && (s[i] < 'a' || 'z' < s[i]) {
			fmt.Println("No")
			return
		} else if i%2 == 1 && (s[i] < 'A' || 'Z' < s[i]) {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

// var sc = bufio.NewScanner((os.Stdin))

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
