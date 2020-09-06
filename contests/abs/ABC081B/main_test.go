package main

import (
	"bytes"
	"testing"
)

func TestAbc081b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"最大で2回操作を行うことができる",
			bytes.NewBufferString("3\n8 12 40"),
			2,
		},
		{
			"最初から奇数5が書かれているため、1回も操作を行うことができない",
			bytes.NewBufferString("4\n5 6 8 10"),
			0,
		},
		{
			"最大で8回操作を行うことができる",
			bytes.NewBufferString("6\n382253568 723152896 37802240 379425024 404894720 471526144"),
			8,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc081b(test.in); got != test.want {
				t.Errorf("abc081b(): got %v want %v", got, test.want)
			}
		})
	}
}
