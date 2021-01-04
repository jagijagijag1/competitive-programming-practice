package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// sc.Split(bufio.ScanWords)
	s, res := nextLine(), []rune{}
	for _, c := range s {
		if c == 'B' && len(res) != 0 {
			res = res[:len(res)-1]
		} else if c == '0' || c == '1' {
			res = append(res, c)
		}
	}
	fmt.Println(string(res))
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

// func nextInt() int {
// 	l := nextLine()
// 	i, e := strconv.Atoi(l)
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }
