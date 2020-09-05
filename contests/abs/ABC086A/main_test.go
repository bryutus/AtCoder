package main

import (
	"testing"
)

func TestAbc086a(t *testing.T) {

	testCases := []struct {
		desc string
		a    int
		b    int
		want string
	}{
		{"3×4=12は偶数なのでEvenを出力", 3, 4, "Even"},
		{"1×21=21は奇数なのでOddを出力", 1, 21, "Odd"},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc086a(test.a, test.b); got != test.want {
				t.Errorf("main(): got %v want %v", got, test.want)
			}
		})
	}
}
