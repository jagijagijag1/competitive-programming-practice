package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC122B() {
	sc.Split(bufio.ScanWords)
	S := nextLine()
	fmt.Printf("%d\n", ABC122B(S))
}

// ABC122B ...
func ABC122B(T string) int {
	cnt, max := 0, 0
	for _, c := range T {
		if c == 'A' || c == 'C' || c == 'G' || c == 'T' {
			cnt++
		} else if cnt != 0 {
			if max < cnt {
				max = cnt
			}
			cnt = 0
		}
	}

	if cnt != 0 && max < cnt {
		max = cnt
	}

	return max
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
