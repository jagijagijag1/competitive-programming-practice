package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc.Split(bufio.ScanWords)
	fmt.Println(strings.Split(nextLine(), ".")[0])
}

var sc = bufio.NewScanner((os.Stdin))

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
