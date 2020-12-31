package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
func mainKeyence2019B() {
	sc.Split(bufio.ScanWords)
	S := nextLine()
	fmt.Printf("%s\n", Keyence2019B(S))
}

// Keyence2019B ...
func Keyence2019B(S string) string {
	for i := 0; i < len(S); i++ {
		for j := i; j <= len(S); j++ {
			if string(S[:i])+string(S[j:]) == "keyence" {
				return "YES"
			}
		}
	}

	return "NO"
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
