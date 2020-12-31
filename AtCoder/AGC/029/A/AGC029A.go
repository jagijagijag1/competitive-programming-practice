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
func mainAGC029A() {
	//sc.Split(bufio.ScanWords)
	S := readLine()

	fmt.Printf("%d\n", AGC029A(S))
}

// AGC029A ...
func AGC029A(S string) int {
	res := 0

	tail := len(S) - 1
	for i := 0; i < len(S); i++ {
		if S[i] == 'B' {
			res += tail - i
			tail--
		}
	}

	return res
}
