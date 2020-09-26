package main

import (
	"bytes"
	"testing"
)

func TestAbc173c(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"cupofcoffeeを1文字ずつ別の文字に書き換えてcupofhotteaにするには4回の操作が必要",
			bytes.NewBufferString(`cupofcoffee
			cupofhottea`),
			4,
		},
		{
			"abcdeを1文字ずつ別の文字に書き換えてbcdeaにするには5回の操作が必要",
			bytes.NewBufferString(`abcde
			bcdea`),
			5,
		},
		{
			"appleを1文字ずつ別の文字に書き換えてappleにするには操作が不要",
			bytes.NewBufferString(`apple
			apple`),
			0,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc172b(test.in); got != test.want {
				t.Errorf("abc172b(): got %v want %v", got, test.want)
			}
		})
	}
}
