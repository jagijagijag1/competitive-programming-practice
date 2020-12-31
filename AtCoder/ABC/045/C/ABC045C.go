package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

// func main() {
func mainABC045C() {
	// sc.Split(bufio.ScanWords)
	S := nextLine()

	fmt.Printf("%d\n", ABC045C(S))
}

// ABC045C ...
func ABC045C(S string) (res int) {
	for bit := 0; bit < (1 << uint(len(S)-1)); bit++ {
		tmp, last := 0, 0
		//tmpstr := ""
		for i := 0; i < len(S)-1; i++ {
			if bit>>uint(i)&1 == 1 {
				num, _ := strconv.Atoi(S[last : i+1])
				tmp += num
				//tmpstr += S[last:i+1] + "+"
				last = i + 1
			}
		}
		num, _ := strconv.Atoi(S[last:])
		tmp += num
		//tmpstr += S[last:]
		//fmt.Println(tmpstr)

		res += tmp
	}

	return
}
