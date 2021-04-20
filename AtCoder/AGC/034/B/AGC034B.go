package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	s, sd, res := []rune(nextLine()), []rune{}, 0

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && s[i] == 'B' && s[i+1] == 'C' {
			sd = append(sd, 'D')
			i++
		} else {
			sd = append(sd, s[i])
		}
	}

	acnt := 0
	for i := 0; i < len(sd); i++ {
		if sd[i] == 'A' {
			acnt++
		} else if sd[i] == 'D' {
			res += acnt
		} else {
			acnt = 0
		}
	}
	fmt.Println(res)
}

// var sc = bufio.NewScanner((os.Stdin))

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
