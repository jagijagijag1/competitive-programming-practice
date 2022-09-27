package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	if n%2 == 1 {
		return
	}
	for bit := 0; bit < (1 << uint(n)); bit++ {
		s := ""
		for i := n - 1; 0 <= i; i-- {
			if bit>>uint(i)&1 == 0 {
				s += "("
			} else {
				s += ")"
			}
		}
		c, cl, cr := 0, 0, 0
		for i := range s {
			if s[i] == '(' {
				c++
				cl++
			} else {
				c--
				cr++
				if c < 0 {
					break
				}
			}
		}
		if c == 0 && cl == cr {
			fmt.Println(s)
		}
	}
}

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

// var sc = bufio.NewScanner((os.Stdin))
