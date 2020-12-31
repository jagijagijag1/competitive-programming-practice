package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC063B() {
	S := nextLine()
	fmt.Printf("%s\n", ABC063B(S))
}

// ABC063B ...
func ABC063B(S string) string {
	used := make([]bool, 26)
	for _, s := range S {
		if used[s-'a'] {
			return "no"
		}
		used[s-'a'] = true
	}

	return "yes"
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
