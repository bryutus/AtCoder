package main

import (
	"bytes"
	"testing"
)

func TestAbc088b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"自分の得点を最大化するように最適な戦略を取った時、先行は後攻より2点多く取る",
			bytes.NewBufferString("2\n3 1"),
			2,
		},
		{
			"自分の得点を最大化するように最適な戦略を取った時、先行は後攻より5点多く取る",
			bytes.NewBufferString("3\n2 7 4"),
			5,
		},
		{
			"自分の得点を最大化するように最適な戦略を取った時、先行は後攻より18点多く取る",
			bytes.NewBufferString("4\n20 18 2 18"),
			18,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc088b(test.in); got != test.want {
				t.Errorf("abc088b(): got %v want %v", got, test.want)
			}
		})
	}
}
