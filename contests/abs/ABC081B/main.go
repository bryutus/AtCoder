package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println(abc081b(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}

	return i
}

func abc081b(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	min := math.MaxInt64
	for i := 0; i < n; i++ {
		x := nextInt(sc)
		cnt := shift(x, 0)
		if cnt < min {
			min = cnt
		}
	}

	return min
}

func shift(x, c int) int {
	if x%2 == 0 {
		return shift(x/2, c+1)
	}

	return c
}
