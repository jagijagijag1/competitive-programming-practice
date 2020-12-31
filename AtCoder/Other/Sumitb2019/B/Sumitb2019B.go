package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// func main() {
func mainSumitb2019B() {
	sc.Split(bufio.ScanWords)
	N := nextInt()

	fmt.Printf("%s\n", Sumitb2019B(N))
}

// Sumitb2019B ...
func Sumitb2019B(N int) string {
	p := math.Floor(float64(N) / 1.08)

	if int(math.Floor(p*1.08)) == N {
		return strconv.Itoa(int(p))
	}
	if int(math.Floor((p+1)*1.08)) == N {
		return strconv.Itoa(int(p + 1))
	}

	return ":("
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
