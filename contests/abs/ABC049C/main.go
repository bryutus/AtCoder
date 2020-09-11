package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(abc049c(os.Stdin))
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func abc049c(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	s := nextString(sc)

	s = strings.Replace(s, "eraser", "0", -1)
	s = strings.Replace(s, "erase", "0", -1)
	s = strings.Replace(s, "dreamer", "0", -1)
	s = strings.Replace(s, "dream", "0", -1)
	s = strings.Replace(s, "0", "", -1)
	s = strings.TrimSpace(s)

	if s == "" {
		return "YES"
	}
	return "NO"
}
