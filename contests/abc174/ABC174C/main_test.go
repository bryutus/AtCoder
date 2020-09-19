package main

import (
	"bytes"
	"testing"
)

func TestAbc174c(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"7,77,777...という数列の中に101の倍数が登場するのは4項目",
			bytes.NewBufferString("101"),
			4,
		},
		{
			"7,77,777...という数列の中に2の倍数は登場しない",
			bytes.NewBufferString("2"),
			-1,
		},
		{
			"7,77,777...という数列の中に999983の倍数が登場するのは999982項目",
			bytes.NewBufferString("999983"),
			999982,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc174c(test.in); got != test.want {
				t.Errorf("abc174c(): got %v want %v", got, test.want)
			}
		})
	}
}
