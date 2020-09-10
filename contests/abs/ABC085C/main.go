package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Println(abc085c(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}

	return i
}

func abc085c(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(sc), nextInt(sc)

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j
			if m == i*10000+j*5000+k*1000 {
				return fmt.Sprintf("%d %d %d", i, j, k)
			}
		}
	}

	return "-1 -1 -1"
}
