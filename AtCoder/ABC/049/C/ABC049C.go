package main

import (
	"fmt"
	"regexp"
	"strings"
)

// var sc = bufio.NewScanner((os.Stdin))
//
// func nextLine() string {
// 	sc.Scan()
// 	return sc.Text()
// }
//
// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

//func main() {
func mainABC049C() {
	//sc.Split(bufio.ScanWords)
	//S := nextLine()
	//S := strings.Trim(nextLine(), "\n")

	var S string
	fmt.Scanf("%s", &S)

	fmt.Printf("%s\n", ABC049C(S))
}

// ABC049C ...
func ABC049C(S string) string {
	r := regexp.MustCompile(`^(dream(er)?|eraser?)*$`)
	//r := regexp.MustCompile(`(dream|dreamer|erase|eraser)*`)
	if r.MatchString(S) {
		return "YES"
	}
	return "NO"
}

// ABC049CWA ...
func ABC049CWA(S string) string {
	unit := []string{"dream", "dreamer", "erase", "eraser"}
	s, t := stack{}, ""
	s = s.Push(S)

	for len(s) != 0 {
		s, t = s.Pop()
		for _, u := range unit {
			if strings.HasPrefix(t, u) {
				tt := t[len(u):]
				if len(tt) == 0 {
					return "YES"
				}

				s = s.Push(tt)
			}
		}
	}

	return "NO"
}
