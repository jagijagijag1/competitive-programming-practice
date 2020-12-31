package main

import (
	"bufio"
	"fmt"
	"os"
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

var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

// func main() {
func mainABC098C() {
	//sc.Split(bufio.ScanWords)
	_, S := readLine(), readLine()

	fmt.Printf("%d\n", ABC098C(S))
}

// ABC098C ...
func ABC098C(S string) int {
	w, e := make([]int, len(S)), make([]int, len(S))

	for i := 1; i < len(S); i++ {
		if S[i-1] == 'W' {
			w[i] = w[i-1] + 1
		} else {
			w[i] = w[i-1]
		}
	}

	for i := len(S) - 2; i >= 0; i-- {
		if S[i+1] == 'E' {
			e[i] = e[i+1] + 1
		} else {
			e[i] = e[i+1]
		}
	}

	res := len(S)
	for leader := 0; leader < len(S); leader++ {
		tmp := w[leader] + e[leader]
		if tmp < res {
			res = tmp
		}
	}

	return res
}
