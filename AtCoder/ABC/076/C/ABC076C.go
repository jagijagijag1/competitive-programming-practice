package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
func mainABC076C() {
	//sc.Split(bufio.ScanWords)
	S, T := nextLine(), nextLine()

	fmt.Printf("%s\n", ABC076C(S, T))
}

// ABC076C ...
func ABC076C(S, T string) string {
	S = strings.ReplaceAll(S, "?", ".")
	for i := len(S) - len(T); i >= 0; i-- {
		h, s, t := S[:i], S[i:i+len(T)], S[i+len(T):]
		matched, _ := regexp.Match(s, []byte(T))
		if matched {
			return h + T + t
		}
		// tmpS := S[i : i+len(T)]
		// r := regexp.MustCompile("^" + tmpS + "$")
		// if r.MatchString(T) {
		// 	li := strings.LastIndex(S, tmpS)
		// 	return strings.ReplaceAll(S[:li]+T+S[li+len(T):], ".", "a")
		// }
	}

	return "UNRESTORABLE"
}

// ABC076Cwa ...
func ABC076Cwa(S, T string) string {
	for i := 0; i < len(S); i++ {
		for j := i; j < len(S); j++ {
			tmpS := S[i : j+1]
			r := regexp.MustCompile("^" + strings.ReplaceAll(tmpS, "?", ".") + "$")
			if r.MatchString(T) {
				res := strings.Replace(S, tmpS, T, 1)
				res = strings.ReplaceAll(res, "?", "a")
				return res
			}
		}
	}

	return "UNRESTORABLE"
}
