package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainPanasonic2020D() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	// fmt.Printf("%d\n", Panasonic2020D(N))
	for _, s := range Panasonic2020D(N) {
		fmt.Printf("%s\n", s)
	}
}

var res []string

// Panasonic2020D ...
func Panasonic2020D(N int) []string {
	res = []string{}
	dfsPanasonic2020D(N, 'a'-1, "")

	return res
}

func dfsPanasonic2020D(N int, maxchar rune, curr string) {
	if len(curr) == N {
		res = append(res, curr)
		return
	}

	for c := 'a'; c <= maxchar+1; c++ {
		tmpstr := curr + string(c)
		if c < maxchar {
			dfsPanasonic2020D(N, maxchar, tmpstr)
		} else {
			dfsPanasonic2020D(N, c, tmpstr)
		}
	}
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
