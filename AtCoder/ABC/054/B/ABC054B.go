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
func mainABC054B() {
	sc.Split(bufio.ScanWords)
	N, M, A, B := nextInt(), nextInt(), []string{}, []string{}

	for i := 0; i < N; i++ {
		A = append(A, nextLine())
	}
	for i := 0; i < M; i++ {
		B = append(B, nextLine())
	}

	fmt.Printf("%s\n", ABC054B(A, B))
}

// ABC054B ...
func ABC054B(A, B []string) string {
	if len(A) == len(B) {
		if checkUnderRect(A, B, 0, 0) {
			return "Yes"
		}
	} else {
		for i := 0; i < len(A)-len(B); i++ {
			for j := 0; j < len(A[i])-len(B); j++ {
				if A[i][j:j+len(B)] == B[0] {
					if checkUnderRect(A, B, i, j) {
						return "Yes"
					}
				}
			}
		}
	}

	return "No"
}

func checkUnderRect(A, B []string, i, j int) bool {
	for k := 1; k < len(B); k++ {
		ak := k + i
		if A[ak][j:j+len(B)] != B[k] {
			return false
		}
	}

	return true
}
