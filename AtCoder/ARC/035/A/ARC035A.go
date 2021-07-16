package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s := nextLine()

	var l, r string
	if len(s)%2 == 0 {
		l, r = s[:len(s)/2], s[len(s)/2:]
	} else {
		l, r = s[:len(s)/2], s[len(s)/2+1:]
	}

	for i := range l {
		if l[i] != r[len(r)-1-i] && l[i] != '*' && r[len(r)-1-i] != '*' {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

var sc = bufio.NewScanner((os.Stdin))

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
