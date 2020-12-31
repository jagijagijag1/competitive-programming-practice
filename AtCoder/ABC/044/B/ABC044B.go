package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC044B() {
	// sc.Split(bufio.ScanWords)
	w := nextLine()

	fmt.Printf("%s\n", ABC044B(w))
}

// ABC044B ...
func ABC044B(w string) string {
	cnt := make([]int, 26)
	for _, c := range w {
		cnt[c-'a']++
	}

	for _, c := range cnt {
		if c%2 == 1 {
			return "No"
		}
	}

	return "Yes"
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
