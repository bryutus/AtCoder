package main

import (
	"bytes"
	"testing"
)

func TestAbc173a(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"1000円札のみを使って支払いを行う時、お釣りは100円",
			bytes.NewBufferString("1900"),
			100,
		},
		{
			"1000円札のみを使って支払いを行う時、お釣りは0円",
			bytes.NewBufferString("3000"),
			0,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc173a(test.in); got != test.want {
				t.Errorf("abc173a(): got %v want %v", got, test.want)
			}
		})
	}
}
