package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// sc.Split(bufio.ScanWords)
	// s := nextLine()
	s := readLine()
	cnt := map[byte]int{}

	for i := range s {
		cnt[s[i]]++
	}

	if len(s) == 1 {
		fmt.Println("YES")
	} else if abs(cnt['a']-cnt['b']) >= 2 || abs(cnt['b']-cnt['c']) >= 2 || abs(cnt['c']-cnt['a']) >= 2 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

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
