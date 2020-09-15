package main

import (
	"bytes"
	"testing"
)

func TestAbc178b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"1 <= x <= 2, 1 <= y <= 1を満たす整数x,yについて、x×yの最大値は2",
			bytes.NewBufferString("1 2 1 1"),
			2,
		},
		{
			"3 <= x <= 5, -4 <= y <= -2を満たす整数x,yについて、x×yの最大値は2",
			bytes.NewBufferString("3 5 -4 -2"),
			-6,
		},
		{
			"-1000000000 <= x <= 0, -1000000000 <= y <= 0を満たす整数x,yについて、x×yの最大値は2",
			bytes.NewBufferString("-1000000000 0 -1000000000 0"),
			1000000000000000000,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc178b(test.in); got != test.want {
				t.Errorf("abc178b(): got %v want %v", got, test.want)
			}
		})
	}
}
