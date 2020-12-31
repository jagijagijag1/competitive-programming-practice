package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC142D() {
	sc.Split(bufio.ScanWords)
	A, B := nextUint64(), nextUint64()

	fmt.Printf("%d\n", ABC142D(A, B))
}

// ABC142D ...
func ABC142D(A, B uint64) int {
	// list common divisor
	cd := []uint64{}
	for i := uint64(1); i*i <= B; i++ {
		if B%i == 0 && A%i == 0 {
			cd = append(cd, i)
		}
		if B/i != i && B%(B/i) == 0 && A%(B/i) == 0 {
			cd = append(cd, B/i)
		}
	}

	// select primes
	p := []uint64{}
	for _, t := range cd {
		isprime := true
		for i := uint64(2); i*i <= t; i++ {
			if t%i == 0 {
				isprime = false
				break
			}
		}

		if isprime {
			p = append(p, t)
		}
	}

	return len(p)
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
