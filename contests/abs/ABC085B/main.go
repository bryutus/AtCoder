package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(abc085b(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}

	return i
}

func abc085b(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt(sc))
	}

	return len(unique(a))
}

func unique(a []int) (unique []int) {
	m := map[int]bool{}

	for _, v := range a {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}
