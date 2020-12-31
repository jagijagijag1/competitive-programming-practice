package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, s := nextInt(), map[string]int{}
	for i := 0; i < N; i++ {
		s[nextLine()]++
	}
	M, t := nextInt(), map[string]int{}
	for i := 0; i < M; i++ {
		t[nextLine()]++
	}

	res := 0
	for k := range s {
		if res < s[k]-t[k] {
			res = s[k] - t[k]
		}
	}

	fmt.Println(res)
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
