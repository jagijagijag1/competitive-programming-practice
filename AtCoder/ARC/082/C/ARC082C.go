package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainARC082C() {
	sc.Split(bufio.ScanWords)
	N, a := nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	fmt.Printf("%d\n", ARC082C(a))
}

// ARC082C ...
func ARC082C(a []int) (res int) {
	cnt := map[int]int{}
	for _, t := range a {
		cnt[t]++
		cnt[t-1]++
		cnt[t+1]++
	}

	for _, v := range cnt {
		if res < v {
			res = v
		}
	}

	return
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
