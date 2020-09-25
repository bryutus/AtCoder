package main

import (
	"bytes"
	"testing"
)

func TestAbc172a(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"2+2^2+2^3=14",
			bytes.NewBufferString(`2`),
			14,
		},
		{
			"10+10^2+10^3=1110",
			bytes.NewBufferString(`10`),
			1110,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc172a(test.in); got != test.want {
				t.Errorf("abc172a(): got %v want %v", got, test.want)
			}
		})
	}
}
