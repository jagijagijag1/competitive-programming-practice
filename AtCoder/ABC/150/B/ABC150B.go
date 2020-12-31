package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC150B() {
	sc.Split(bufio.ScanWords)
	_, S := nextInt(), nextLine()

	fmt.Printf("%d\n", ABC150B(S))
}

// ABC150B ...
func ABC150B(S string) (res int) {
	for i := 0; i < len(S)-2; i++ {
		if S[i] == 'A' && S[i:i+3] == "ABC" {
			res++
			i += 2
			if len(S) <= i {
				break
			}
		} else {
			continue
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
