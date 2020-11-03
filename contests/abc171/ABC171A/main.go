package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	fmt.Println(abc171a(os.Stdin))
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func abc171a(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s := nextString(sc)
	b := regexp.MustCompile(`[A-Z]`).Match([]byte(s))
	if b == true {
		return "A"
	}

	return "a"
}
