package main

import (
	"bytes"
	"testing"
)

func TestAbc178c(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"",
			bytes.NewBufferString("2"),
			2,
		},
		{
			"",
			bytes.NewBufferString("1"),
			0,
		},
		{
			"",
			bytes.NewBufferString("869121"),
			2511445,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc178c(test.in); got != test.want {
				t.Errorf("abc178c(): got %v want %v", got, test.want)
			}
		})
	}
}
