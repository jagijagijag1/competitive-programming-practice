package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextLine()
	if strings.Contains(n, "7") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
