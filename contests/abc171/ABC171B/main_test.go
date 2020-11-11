package main

import (
	"bytes"
	"testing"
)

func TestAbc171b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"5種類の果物から3種類を買うときの最小合計価格は210円",
			bytes.NewBufferString(`5 3
			50 100 80 120 80`),
			210,
		},
		{
			"1種類の果物から1種類を買うときの最小合計価格は1000円",
			bytes.NewBufferString(`1 1
			1000`),
			1000,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc171b(test.in); got != test.want {
				t.Errorf("abc171b(): got %v want %v", got, test.want)
			}
		})
	}
}
