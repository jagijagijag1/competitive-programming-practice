package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
func mainARC065A() {
	// sc.Split(bufio.ScanWords)
	S := readLine()

	fmt.Printf("%s\n", ARC065A(S))
}

// ARC065A ...
func ARC065A(S string) string {
	r := regexp.MustCompile(`^(dream|dreamer|erase|eraser)+$`)

	if r.MatchString(S) {
		return "YES"
	}

	return "NO"
}
