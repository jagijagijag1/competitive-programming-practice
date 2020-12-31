package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, m := nextInt(), map[int]struct{}{}
	for i := 0; i < N; i++ {
		a := nextInt()
		if _, ok := m[a]; ok {
			delete(m, a)
		} else {
			m[a] = struct{}{}
		}
	}

	fmt.Println(len(m))
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
