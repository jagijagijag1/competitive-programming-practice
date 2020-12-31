package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, s, t := nextInt(), nextLine(), nextLine()

	res := s + t
	for i := N - 1; 0 <= i; i-- {
		if s[N-1-i:] == t[0:i+1] {
			res = s + t[i+1:]
			break
		}
	}

	fmt.Println(len(res))
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
