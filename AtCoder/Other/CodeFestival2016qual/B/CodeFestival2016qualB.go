package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainCodeFestival2016qualB() {
	sc.Split(bufio.ScanWords)
	N, a := nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	fmt.Printf("%d\n", CodeFestival2016qualB(a))
}

// CodeFestival2016qualB ...
func CodeFestival2016qualB(a []int) (res int) {
	for i, aa := range a {
		if a[aa-1] == i+1 {
			res++
		}
	}

	return res / 2
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
