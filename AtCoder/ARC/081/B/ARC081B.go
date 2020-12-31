package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

// func main() {
func mainARC081B() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
	S1, S2 := nextLine(), nextLine()

	fmt.Printf("%d\n", ARC081B(N, S1, S2))
}

// ARC081B ...
func ARC081B(N int, S1, S2 string) (res int) {
	mod := 1000000007
	next, last := 0, ""

	// check first block(s)
	if S1[0] == S2[0] {
		// if the first block is vertical
		res += 3
		next++
		last = "v"
	} else {
		// if the first blocks are horizontal
		res += 6
		next = 2
		last = "h"
	}

	for i := next; i < N; {
		if S1[i] == S2[i] {
			if last == "v" {
				// v -> v
				res = (res * 2) % mod
			} else {
				// h -> v
				res = (res * 1) % mod
			}
			last = "v"
			i++
		} else {
			if last == "v" {
				// v -> h
				res = (res * 2) % mod
			} else {
				// h -> h
				res = (res * 3) % mod
			}
			last = "h"
			i += 2
		}
	}

	return
}
