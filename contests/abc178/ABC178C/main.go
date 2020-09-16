package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const mod = 1000000007

func main() {
	fmt.Println(abc178c(os.Stdin))
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func abc178c(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)

	ans := (pow(10, n) - (pow(9, n))*2 + pow(8, n)) % mod

	return (ans + mod) % mod
}

func pow(base, exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= base
		res %= mod
	}
	return res
}
