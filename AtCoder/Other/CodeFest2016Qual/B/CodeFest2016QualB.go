package main

import (
	"fmt"
)

// func main() {
func mainCodeFest2016QualB() {
	var N, A, B int
	var S string
	fmt.Scan(&N, &A, &B, &S)
	// sc.Split(bufio.ScanWords)
	// _, A, B, S := nextInt(), nextInt(), nextInt(), nextLine()

	//fmt.Printf("%s\n", CodeFest2016QualB(A, B, S))
	for _, s := range CodeFest2016QualB(A, B, S) {
		fmt.Printf("%s\n", s)
	}
}

// CodeFest2016QualB ...
func CodeFest2016QualB(A, B int, S string) (res []string) {
	p, bp := 0, 0
	for _, s := range S {
		if s == 'a' {
			if p < A+B {
				res = append(res, "Yes")
				p++
			} else {
				res = append(res, "No")
			}
		} else if s == 'b' {
			if p < A+B && bp < B {
				res = append(res, "Yes")
				p++
				bp++
			} else {
				res = append(res, "No")
			}
		} else if s == 'c' {
			res = append(res, "No")
		}
	}

	return
}

// var sc = bufio.NewScanner((os.Stdin))

// func nextLine() string {
// 	sc.Scan()
// 	return sc.Text()
// }

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }
