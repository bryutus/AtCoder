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
	fmt.Println(abc174b(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func abc174b(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, d := nextInt(sc), float64(nextInt(sc))

	var q int
	for i := 0; i < n; i++ {
		x, y := float64(nextInt(sc)), float64(nextInt(sc))
		if math.Sqrt(x*x+y*y) <= d {
			q++
		}
	}
	return q
}
