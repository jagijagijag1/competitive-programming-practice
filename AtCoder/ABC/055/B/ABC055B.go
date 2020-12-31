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
func mainABC055B() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	fmt.Printf("%d\n", ABC055B(N))
}

// ABC055B ...
func ABC055B(N int) int {
	res := 1
	for i := 1; i <= N; i++ {
		res = (res * i) % (1000000000 + 7)
	}
	return res
}
