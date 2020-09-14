package main

import (
	"bytes"
	"testing"
)

func TestAbc178a(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"xが0なら1を出力する",
			bytes.NewBufferString("0"),
			1,
		},
		{
			"xが1なら0を出力する",
			bytes.NewBufferString("1"),
			0,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc178a(test.in); got != test.want {
				t.Errorf("abc178a(): got %v want %v", got, test.want)
			}
		})
	}
}
