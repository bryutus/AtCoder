package main

import (
	"testing"
)

func TestAbc081a(t *testing.T) {

	testCases := []struct {
		desc string
		in   string
		want int
	}{
		{"マス1,3にビー玉が置かれる", "101", 2},
		{"マス3にビー玉が置かれる", "001", 1},
		{"ビー玉はどのマスにも置かれない", "000", 0},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc081a(test.in); got != test.want {
				t.Errorf("main(): got %v want %v", got, test.want)
			}
		})
	}
}
