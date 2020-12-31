package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
func mainABC114C() {
	sc.Split(bufio.ScanWords)
	N := nextLine()

	fmt.Printf("%d\n", ABC114C(N))
}

var resABC114C int

// ABC114C ...
func ABC114C(N string) int {
	dfsABC114C("", N)
	return resABC114C
}

func dfsABC114C(current, upper string) {
	c, _ := strconv.Atoi(current)
	if len(current) == len(upper) {
		u, _ := strconv.Atoi(upper)
		if c <= u && containsAll(current) {
			resABC114C++
		}
		return
	}

	nums := []string{"3", "5", "7"}
	if c == 0 {
		nums = append(nums, "0")
	}
	for _, n := range nums {
		dfsABC114C(current+n, upper)
	}
}

func containsAll(s string) bool {
	if strings.Contains(s, "3") && strings.Contains(s, "5") && strings.Contains(s, "7") {
		return true
	}
	return false
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
