package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, last, t, res := nextInt(), nextInt(), "", 0
	for i := 1; i < N; i++ {
		curr := nextInt()
		if t == "" {
			if last < curr {
				t = "inc"
			} else if curr < last {
				t = "dec"
			}
		} else if t == "inc" && curr < last {
			t = ""
			res++
		} else if t == "dec" && last < curr {
			t = ""
			res++
		}
		last = curr
	}
	fmt.Println(res + 1)
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
