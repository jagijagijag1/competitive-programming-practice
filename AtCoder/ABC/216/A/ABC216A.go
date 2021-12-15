package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	xy := nextLine()
	x, y := xy[:len(xy)-2], int(xy[len(xy)-1]-'0')
	if 0 <= y && y <= 2 {
		fmt.Println(x + "-")
	} else if y <= 6 {
		fmt.Println(x)
	} else {
		fmt.Println(x + "+")
	}
}

type tuple struct {
	i, a int
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
// l := nextLine()
// i, e := strconv.Atoi(l)
// if e != nil {
// panic(e)
// }
// return i
// }

// var sc = bufio.NewScanner((os.Stdin))
