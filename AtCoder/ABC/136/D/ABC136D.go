package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
func mainABC136D() {
	// sc.Split(bufio.ScanWords)
	S := readLine()

	fmt.Printf("%s\n", ABC136D(S))
}

// ABC136D ...
func ABC136D(S string) (resstr string) {
	res := make([]int, len(S))

	succ := 1
	for i := 1; i < len(S); i++ {
		if succ != 0 && S[i] == 'L' {
			if succ%2 == 0 {
				res[i-1] += succ / 2
				res[i] += succ / 2
			} else {
				res[i-1] += succ/2 + 1
				res[i] += succ / 2
			}
			succ = 0
		} else if S[i] == 'R' {
			succ++
		}
	}

	succ = 1
	for i := len(S) - 2; i >= 0; i-- {
		if succ != 0 && S[i] == 'R' {
			if succ%2 == 0 {
				res[i] += succ / 2
				res[i+1] += succ / 2
			} else {
				res[i] += succ / 2
				res[i+1] += succ/2 + 1
			}
			succ = 0
		} else if S[i] == 'L' {
			succ++
		}
	}

	for i := 0; i < len(res)-1; i++ {
		resstr += strconv.Itoa(res[i]) + " "
	}
	resstr += strconv.Itoa(res[len(res)-1])

	return
}

// ABC136Dnaive ...
func ABC136Dnaive(S string) (res string) {
	resSlice := make([]int, len(S))

	for i := 0; i < len(S); i++ {
		pos, cnt := i, 0
		for {
			if S[pos] == 'R' && S[pos+1] == 'L' {
				if cnt%2 == 0 {
					resSlice[pos]++
				} else {
					resSlice[pos+1]++
				}
				break
			} else if S[pos] == 'L' && S[pos-1] == 'R' {
				if cnt%2 == 0 {
					resSlice[pos]++
				} else {
					resSlice[pos-1]++
				}
				break
			} else if S[pos] == 'R' {
				pos++
				cnt++
			} else if S[pos] == 'L' {
				pos--
				cnt++
			}
		}
	}

	for i := 0; i < len(resSlice); i++ {
		res += strconv.Itoa(resSlice[i]) + " "
	}

	return res[0 : len(res)-1]
}
