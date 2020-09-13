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
	fmt.Println(abc086c(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func abc086c(r io.Reader) string {
	var n, t, nt int
	var x, y, nx, ny float64

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n = nextInt(sc)

	for i := 0; i < n; i++ {
		nt, nx, ny = nextInt(sc), float64(nextInt(sc)), float64(nextInt(sc))
		dt := nt - t
		dist := math.Abs(nx-x) + math.Abs(ny-y)
		if dt < int(dist) || int(dist)%2 != dt%2 {

			return "No"
		}
		x, y, t = nx, ny, nt
	}

	return "Yes"
}
