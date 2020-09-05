package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(abc086a(a, b))
}

func abc086a(a, b int) string {
	if a*b%2 == 0 {
		return "Even"
	}
	return "Odd"
}
