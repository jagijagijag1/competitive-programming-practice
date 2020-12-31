package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, a := nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	ci, res := 1, 0
	for len(a) > 0 {
		h := a[0]
		a = a[1:]

		if h == ci {
			ci++
		} else {
			res++
		}
	}

	if res == N {
		fmt.Println(-1)
	} else {
		fmt.Println(res)
	}
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
