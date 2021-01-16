package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	N, P, modcnt := nextInt(), nextInt(), map[int]int{}
	for i := 0; i < N; i++ {
		if nextInt()%2 == 0 {
			modcnt[0]++
		} else {
			modcnt[1]++
		}
	}

	if modcnt[1] == 0 {
		if P == 0 {
			fmt.Println(int(math.Pow(2, float64(N))))
		} else {
			fmt.Println(0)
		}
	} else {
		fmt.Println(int(math.Pow(2, float64(N-1))))
	}

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
