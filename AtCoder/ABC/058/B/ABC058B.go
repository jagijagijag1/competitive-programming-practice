package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC058B() {
	// sc.Split(bufio.ScanWords)
	O, E := nextLine(), nextLine()

	fmt.Printf("%s\n", ABC058B(O, E))
}

// ABC058B ...
func ABC058B(O, E string) string {
	res := []byte{}
	for i := range O {
		if i < len(O) {
			res = append(res, O[i])
		}
		if i < len(E) {
			res = append(res, E[i])
		}
	}

	return string(res)
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
