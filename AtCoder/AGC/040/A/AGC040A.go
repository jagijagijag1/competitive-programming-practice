package main

import (
	"bufio"
	"fmt"
	"math"
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
func mainAGC040A() {
	//sc.Split(bufio.ScanWords)
	S := readLine()

	fmt.Printf("%d\n", AGC040A(S))
}

// AGC040Ashort ...
func AGC040Ashort(S string) int {
	v := make([]int, len(S)+1)

	for i := 0; i < len(S); i++ {
		if S[i] == '<' {
			v[i+1] = v[i] + 1
		}
	}

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '>' {
			v[i] = int(math.Max(float64(v[i]), float64(v[i+1]+1)))
		}
	}

	res := 0
	for i := 0; i < len(S)+1; i++ {
		res += v[i]
	}

	return res
}

// AGC040A ...
func AGC040A(S string) int {
	res := 0
	gtlist, ltlist := make([]int, len(S)+1), make([]int, len(S)+1)
	for i := 0; i < len(S)+1; i++ {
		gtlist[i], ltlist[i] = -1, -1
	}

	contd := false
	for i := 0; i < len(S)+1; i++ {
		if i == len(S) {
			// in case of tail
			if S[i-1] == '<' {
				gtlist[i] = gtlist[i-1] + 1
			}
		} else if S[i] == '<' {
			if contd {
				gtlist[i] = gtlist[i-1] + 1
			} else {
				gtlist[i] = 0
				contd = true
			}
		} else if contd {
			gtlist[i] = gtlist[i-1] + 1
			contd = false
		}
	}

	contd = false
	for i := len(S); i >= 0; i-- {
		if i == 0 {
			// in case of head
			if S[i] == '>' {
				ltlist[i] = ltlist[i+1] + 1
			}
		} else if S[i-1] == '>' {
			if contd {
				ltlist[i] = ltlist[i+1] + 1
			} else {
				ltlist[i] = 0
				contd = true
			}
		} else if contd {
			ltlist[i] = ltlist[i+1] + 1
			contd = false
		}
	}

	// merge
	for i := 0; i < len(S)+1; i++ {
		if gtlist[i] != -1 && ltlist[i] != -1 {
			if gtlist[i] > ltlist[i] {
				res += gtlist[i]
			} else {
				res += ltlist[i]
			}
		} else if gtlist[i] != -1 {
			res += gtlist[i]
		} else if ltlist[i] != -1 {
			res += ltlist[i]
		} else {
			return -1
		}
	}

	return res
}
