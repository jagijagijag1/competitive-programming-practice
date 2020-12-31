package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainABC128C() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()

	s := make([][]int, M)
	for i := 0; i < M; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			s[i] = append(s[i], nextInt())
		}
	}

	p := []int{}
	for i := 0; i < M; i++ {
		p = append(p, nextInt())
	}

	fmt.Printf("%d\n", ABC128C(N, M, s, p))
}

// ABC128C ...
func ABC128C(N, M int, s [][]int, p []int) (res int) {

	for bit := 0; bit < (1 << N); bit++ {
		allOn := true
		for i := range s {
			onSwitch := 0
			for j := range s[i] {
				if bit>>uint(s[i][j]-1)&1 == 1 {
					onSwitch++
				}
			}

			if onSwitch%2 != p[i] {
				allOn = false
				break
			}
		}

		if allOn {
			res++
		}
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
