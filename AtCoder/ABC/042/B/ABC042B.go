package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// func main() {
func mainABC042B() {
	sc.Split(bufio.ScanWords)
	N, _, S := nextInt(), nextInt(), []string{}
	for i := 0; i < N; i++ {
		S = append(S, nextLine())
	}

	fmt.Printf("%s\n", ABC042B(S))
}

// ABC042B ...
func ABC042B(S []string) (res string) {
	sort.Strings(S)
	for _, s := range S {
		res += s
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
