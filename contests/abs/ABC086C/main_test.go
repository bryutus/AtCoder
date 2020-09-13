package main

import (
	"bytes"
	"testing"
)

func TestAbc086c(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want string
	}{
		{
			"(0,0)にいる状態から3秒後に(1,2)、6秒後に(1,1)にいるのは可能",
			bytes.NewBufferString(`2
			3 1 2
			6 1 1`),
			"Yes",
		},
		{
			"(0,0)にいる状態から2秒後に(100,100)にいるのは不可能",
			bytes.NewBufferString(`1
			2 100 100`),
			"No",
		},
		{
			"(0,0)にいる状態から5秒後に(1,1)、100秒後に(1,1)にいるのは不可能",
			bytes.NewBufferString(`2
			5 1 1
			100 1 1`),
			"No",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc086c(test.in); got != test.want {
				t.Errorf("abc086c(): got %v want %v", got, test.want)
			}
		})
	}
}
