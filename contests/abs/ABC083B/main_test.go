package main

import (
	"bytes"
	"testing"
)

func TestAbc083b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"1以上20以下の整数のうち、10進数での各行の和が2以上5以下であるものの総和は84である",
			bytes.NewBufferString("20 2 5"),
			84,
		},
		{
			"1以上10以下の整数のうち、10進数での各行の和が1以上2以下であるものの総和は13である",
			bytes.NewBufferString("10 1 2"),
			13,
		},
		{
			"1以上100以下の整数のうち、10進数での各行の和が4以上16以下であるものの総和は4554である",
			bytes.NewBufferString("100 4 16"),
			4554,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc083b(test.in); got != test.want {
				t.Errorf("abc083b(): got %v want %v", got, test.want)
			}
		})
	}
}
