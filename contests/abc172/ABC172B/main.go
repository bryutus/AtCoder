package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(abc172b(os.Stdin))
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func abc172b(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s, t := nextString(sc), nextString(sc)
	cnt := 0
	len := len(s)
	for i := 0; i < len; i++ {
		if s[i] != t[i] {
			cnt++
		}
	}

	return cnt
}
