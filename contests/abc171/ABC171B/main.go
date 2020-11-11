package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(abc171b(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func abc171b(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n, k := nextInt(sc), nextInt(sc)

	var a []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt(sc))
	}

	sort.Sort(sort.IntSlice(a))

	var t int
	for i := 0; i < k; i++ {
		t += a[i]
	}

	return t
}
