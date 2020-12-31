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

func nextUint64() uint64 {
	l := nextLine()
	i, e := strconv.ParseUint(l, 10, 64)
	if e != nil {
		panic(e)
	}
	return i
}

// func main() {
func mainABC148E() {
	sc.Split(bufio.ScanWords)
	N := nextUint64()

	fmt.Printf("%d\n", ABC148E(N))
}

// ABC148E ...
func ABC148E(N uint64) (res uint64) {
	if N%2 == 1 {
		return
	}

	for i := uint64(10); i <= N; i *= 5 {
		res += N / i
	}

	return
}

// ABC148Enaive ...
func ABC148Enaive(N uint64) (res uint64) {
	sum := uint64(1)
	// restr := ""
	for i := uint64(N); i > 1; i -= 2 {
		sum *= i
		// restr += strconv.Itoa(i) + " * "
	}
	fmt.Println(sum)
	// fmt.Println(restr)

	for {
		if sum%10 != 0 {
			break
		} else {
			res++
			sum = sum / 10
		}
	}

	return
}
