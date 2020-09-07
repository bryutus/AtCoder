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
	fmt.Println(abc088b(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}

	return i
}

func abc088b(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	var a []int
	for i := 0; i < n; i++ {
		a = append(a, nextInt(sc))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	var first, second int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			first += a[i]
		} else {
			second += a[i]
		}
	}

	return first - second
}
