package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC123B() {
	sc.Split(bufio.ScanWords)
	A, B, C, D, E := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	fmt.Printf("%d\n", ABC123B(A, B, C, D, E))
}

// ABC123B ...
func ABC123B(A, B, C, D, E int) (res int) {
	dishes := []int{A, B, C, D, E}
	minmod, minidx := 10, 10
	for i, d := range dishes {
		mod := d % 10
		if mod != 0 && mod < minmod {
			minmod = mod
			minidx = i
		}
	}

	for i, d := range dishes {
		res += d
		mod := d % 10
		if i != minidx && mod != 0 {
			res += 10 - mod
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
