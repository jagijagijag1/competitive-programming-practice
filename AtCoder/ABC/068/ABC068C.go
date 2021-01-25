package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, M, ships := nextInt(), nextInt(), map[int][]int{}
	for i := 0; i < M; i++ {
		a, b := nextInt(), nextInt()
		ships[a] = append(ships[a], b)
	}

	for _, s1 := range ships[1] {
		for _, s2 := range ships[s1] {
			if s2 == N {
				fmt.Println("POSSIBLE")
				return
			}
		}
	}

	fmt.Println("IMPOSSIBLE")
}

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

// var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

// func readLine() string {
// 	buf := make([]byte, 0, 1000000)
// 	for {
// 		l, p, e := rdr.ReadLine()
// 		if e != nil {
// 			panic(e)
// 		}
// 		buf = append(buf, l...)
// 		if !p {
// 			break
// 		}
// 	}
// 	return string(buf)
// }
