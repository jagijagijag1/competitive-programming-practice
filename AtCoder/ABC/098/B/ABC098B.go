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
func mainABC098B() {
	//sc.Split(bufio.ScanWords)
	_, S := nextLine(), nextLine()

	fmt.Printf("%d\n", ABC098B(S))
}

// ABC098B ...
func ABC098B(S string) (res int) {
	alph := "abcdefghijklmnopqrstuvwxyz"
	for i := 1; i < len(S); i++ {
		cnt := 0
		for _, a := range alph {
			l, r := false, false

			for j := 0; j < i; j++ {
				if S[j] == byte(a) {
					l = true
				}
			}
			for j := i; j < len(S); j++ {
				if S[j] == byte(a) {
					r = true
				}
			}

			if l && r {
				cnt++
			}
		}

		if res < cnt {
			res = cnt
		}
	}

	return
}

// ABC098Bnaive ...
func ABC098Bnaive(S string) (res int) {
	for i := 1; i < len(S)-1; i++ {
		chars, checked, cnt := map[string]bool{}, map[string]bool{}, 0
		for j := 0; j < len(S); j++ {
			c := string(S[j])
			if j < i {
				chars[c] = true
			} else {
				_, exists := chars[c]
				_, isChecked := checked[c]
				if exists && !isChecked {
					cnt++
				}
				checked[c] = true
			}
		}

		if res < cnt {
			res = cnt
		}
	}

	return
}
