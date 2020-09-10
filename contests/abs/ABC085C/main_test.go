package main

import (
	"bytes"
	"testing"
)

func TestAbc085c(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want string
	}{
		{
			"9枚のお札の合計金額が45000円となる組み合わせの一つに、10000円札 0枚、5000円札 9枚、1000円札 0枚がある",
			bytes.NewBufferString("9 45000"),
			"0 9 0",
		},
		{
			"20枚のお札の合計金額が196000円となる組み合わせは存在しない",
			bytes.NewBufferString("20 196000"),
			"-1 -1 -1",
		},
		{
			"1000枚のお札の合計金額が1234000円となる組み合わせの一つに、10000円札 14枚、5000円札 27枚、1000円札 959枚がある",
			bytes.NewBufferString("1000 1234000"),
			"2 54 944",
		},
		{
			"2000枚のお札の合計金額が20000000円となる組み合わせの一つに、10000円札 2000枚、5000円札 0枚、1000円札 0枚がある",
			bytes.NewBufferString("2000 20000000"),
			"2000 0 0",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc085c(test.in); got != test.want {
				t.Errorf("abc085c(): got %v want %v", got, test.want)
			}
		})
	}
}
