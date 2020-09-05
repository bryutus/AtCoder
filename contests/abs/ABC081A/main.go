package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(abc081a(s))
}

func abc081a(in string) (c int) {
	for n := range in {
		if in[n] == '1' {
			c++
		}
	}
	return
}
