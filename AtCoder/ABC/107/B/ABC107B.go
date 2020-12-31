package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
func mainABC107B() {
	sc.Split(bufio.ScanWords)
	H, _, a := nextInt(), nextInt(), []string{}
	for i := 0; i < H; i++ {
		a = append(a, nextLine())
	}

	for _, r := range ABC107B(a) {
		fmt.Printf("%s\n", r)
	}
	// fmt.Printf("%d\n", ABC107B())
}

// ABC107B ...
func ABC107B(a []string) (res []string) {
	tmp := []string{}
	for _, s := range a {
		if strings.Contains(s, "#") {
			tmp = append(tmp, s)
		}
	}

	if len(tmp) == 0 {
		return
	}

	flags := []bool{}
	for i := 0; i < len(tmp[0]); i++ {
		flags = append(flags, true)
	}
	for i := 0; i < len(tmp[0]); i++ {
		for _, s := range tmp {
			if s[i] == '#' {
				flags[i] = false
				break
			}
		}
	}

	for _, t := range tmp {
		ress := ""
		for i, tt := range t {
			if !flags[i] {
				ress += string(tt)
			}
		}
		res = append(res, ress)
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
