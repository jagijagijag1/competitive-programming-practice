package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, T, t := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		t = append(t, nextInt())
	}

	res := 0
	for i := 0; i < N-1; i++ {
		if t[i+1] < t[i]+T {
			res += t[i+1] - t[i]
		} else {
			res += T
		}
	}

	fmt.Println(res + T)
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
