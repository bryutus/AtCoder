package main

import (
	"bytes"
	"testing"
)

func TestAbc085b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"重複した値を削除した時の要素数は3である",
			bytes.NewBufferString("4\n10\n8\n8\n6"),
			3,
		},
		{
			"重複した値を削除した時の要素数は1である",
			bytes.NewBufferString("3\n15\n15\n15"),
			1,
		},
		{
			"重複した値を削除した時の要素数は4である",
			bytes.NewBufferString("7\n50\n30\n50\n100\n50\n80\n30"),
			4,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc085b(test.in); got != test.want {
				t.Errorf("abc085b(): got %v want %v", got, test.want)
			}
		})
	}
}
