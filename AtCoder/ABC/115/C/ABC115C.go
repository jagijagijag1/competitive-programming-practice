package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, K, h := nextInt(), nextInt(), []int{}
	for i := 0; i < N; i++ {
		h = append(h, nextInt())
	}

	sort.Ints(h)
	res := []int{}
	for i := 0; i <= N-K; i++ {
		res = append(res, h[i+K-1]-h[i])
	}
	sort.Ints(res)
	fmt.Println(res[0])
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
