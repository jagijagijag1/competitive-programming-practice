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

//func main() {
func mainABC085B() {
	sc.Split(bufio.ScanWords)
	N, a := nextInt(), []int{}
	for i := 0; i < N; i++ {
		a = append(a, nextInt())
	}

	fmt.Printf("%d\n", ABC085B(a))
}

// ABC085B ...
func ABC085B(a []int) int {
	return len(intSliceUnique(a))
}

func intSliceUnique(target []int) (unique []int) {
	m := map[int]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
