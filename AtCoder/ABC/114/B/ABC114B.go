package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC114B() {
	// sc.Split(bufio.ScanWords)
	S := nextLine()

	fmt.Printf("%d\n", ABC114B(S))
}

// ABC114B ...
func ABC114B(S string) int {
	base := 753

	res := 100000000
	for i := 0; i < len(S)-2; i++ {
		tmp, _ := strconv.Atoi(string(S[i : i+3]))
		res = minInt(res, absInt(base-tmp))
	}

	return res
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
