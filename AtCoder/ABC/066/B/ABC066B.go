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

// func main() {
func mainABC066B() {
	//sc.Split(bufio.ScanWords)
	S := nextLine()

	fmt.Printf("%d\n", ABC066B(S))
}

// ABC066B ...
func ABC066B(S string) int {
	for i := len(S) - 2; i > 0; i-- {
		if i%2 == 0 {
			continue
		}

		sl, sr := S[0:i/2+1], S[i/2+1:i+1]
		if sl == sr {
			return i + 1
		}
	}

	return -1
}
