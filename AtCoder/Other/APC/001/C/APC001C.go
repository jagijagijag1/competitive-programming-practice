package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	s := make([]string, N)

	var ans string
	fmt.Println(0)
	fmt.Scan(&ans)
	if ans == "Vacant" {
		return
	}
	s[0] = ans

	fmt.Println(N - 1)
	fmt.Scan(&ans)
	if ans == "Vacant" {
		return
	}
	s[N-1] = ans

	l, r := 0, N-1
	for i := 0; i < 20; i++ {
		m := (l + r) / 2
		fmt.Println(m)
		fmt.Scan(&ans)
		if ans == "Vacant" {
			return
		}

		s[m] = ans
		//if (m-l)%2 == 1 && s[l] == s[m] {
		if ((m-l)%2 == 1 && s[l] == s[m]) || ((m-l)%2 == 0 && s[l] != s[m]) {
			r = m
		} else {
			l = m
		}
	}
}
