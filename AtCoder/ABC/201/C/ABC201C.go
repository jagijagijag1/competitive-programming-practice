package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	s := nextLine()
	in, ex := []int{}, []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'o' {
			in = append(in, i)
		}
		if s[i] == 'x' {
			ex = append(ex, i)
		}
	}

	res := 0
	for i := 0; i < 10; i++ {
		if contains(ex, i) {
			continue
		}
		for j := 0; j < 10; j++ {
			if contains(ex, j) {
				continue
			}
			for k := 0; k < 10; k++ {
				if contains(ex, k) {
					continue
				}
				for l := 0; l < 10; l++ {
					if contains(ex, l) {
						continue
					}
					pass, all := []int{i, j, k, l}, true
					for _, v := range in {
						if !contains(pass, v) {
							all = false
							break
						}
					}
					if all {
						res++
					}
				}
			}
		}
	}
	fmt.Println(res)
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

const (
	initialBufSize = 10000
	maxBufSize     = 1000000
)

var (
	sc *bufio.Scanner = func() *bufio.Scanner {
		sc := bufio.NewScanner(os.Stdin)
		buf := make([]byte, initialBufSize)
		sc.Buffer(buf, maxBufSize)
		return sc
	}()
)
