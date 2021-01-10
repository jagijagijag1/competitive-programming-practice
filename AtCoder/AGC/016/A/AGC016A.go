package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// sc.Split(bufio.ScanWords)
	s := []rune(nextLine())
	res := 100000000

	for i := 'a'; i <= 'z'; i++ {
		ss, cnt, target := s, 0, strings.Repeat(string(i), len(s))
		for string(ss) != target[0:len(ss)] {
			sss := []rune{}
			for j := 0; j < len(ss)-1; j++ {
				if ss[j] == i || ss[j+1] == i {
					sss = append(sss, i)
				} else {
					sss = append(sss, ss[j])
				}
			}
			ss = sss
			cnt++
		}

		if cnt < res {
			res = cnt
		}
	}

	fmt.Println(res)
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
