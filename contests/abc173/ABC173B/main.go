package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(abc173b(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func abc173b(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)

	var ac, wa, tle, re int
	for i := 0; i < n; i++ {
		s := nextString(sc)
		switch s {
		case "AC":
			ac++
		case "WA":
			wa++
		case "TLE":
			tle++
		case "RE":
			re++
		default:
		}
	}

	return fmt.Sprintf("AC x %d\nWA x %d\nTLE x %d\nRE x %d", ac, wa, tle, re)
}
