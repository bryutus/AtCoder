package main

import (
	"bytes"
	"testing"
)

func TestAbc173b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want string
	}{
		{
			"ジャッジ結果はそれぞれ個数が、AC:3, WA:1, TLE:2, RE:0になる。",
			bytes.NewBufferString(`6
			AC
			TLE
			AC
			AC
			WA
			TLE`),
			`AC x 3
WA x 1
TLE x 2
RE x 0`,
		},
		{
			"ジャッジ結果はそれぞれ個数が、AC:10, WA:0, TLE:0, RE:0になる。",
			bytes.NewBufferString(`10
			AC
			AC
			AC
			AC
			AC
			AC
			AC
			AC
			AC
			AC`),
			`AC x 10
WA x 0
TLE x 0
RE x 0`,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc173b(test.in); got != test.want {
				t.Errorf("abc173b(): got %v want %v", got, test.want)
			}
		})
	}
}
