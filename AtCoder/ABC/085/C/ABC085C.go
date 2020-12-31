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
func mainABC085C() {
	sc.Split(bufio.ScanWords)
	N, Y := nextInt(), nextInt()

	fmt.Printf("%s\n", ABC085C(N, Y))
}

// ABC085C ...
func ABC085C(N, Y int) string {
	for i := 0; i <= Y/10000; i++ {
		for j := 0; j <= Y/5000 && i+j <= N; j++ {
			k := N - i - j
			y := i*10000 + j*5000 + k*1000
			if y == Y {
				return strconv.Itoa(i) + " " + strconv.Itoa(j) + " " + strconv.Itoa(k)
			}
		}
	}

	return "-1 -1 -1"
}
