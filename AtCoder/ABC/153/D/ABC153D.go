package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainABC153D() {
	sc.Split(bufio.ScanWords)
	H := nextUint64()

	fmt.Printf("%d\n", ABC153D(H))
}

// ABC153D ...
func ABC153D(H uint64) (res uint64) {
	for i := 0; H != 0; i++ {
		H = H / 2
		res += uint64(math.Pow(float64(2), float64(i)))
	}

	return
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

func nextUint64() uint64 {
	l := nextLine()
	i, e := strconv.ParseUint(l, 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}
