package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(abc083b(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}

	return i
}

func abc083b(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, a, b := nextInt(sc), nextInt(sc), nextInt(sc)
	cnt := 0
	for i := 1; i <= n; i++ {
		summed := sum(i)
		if a <= summed && summed <= b {
			cnt = cnt + i
		}
	}
	return cnt
}

func sum(n int) int {
	if n < 10 {
		return n
	}
	return sum(n/10) + n%10
}
