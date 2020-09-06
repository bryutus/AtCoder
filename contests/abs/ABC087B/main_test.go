package main

import (
	"bytes"
	"testing"
)

func TestAbc087b(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want int
	}{
		{
			"500円玉:2枚、100円玉:2枚、50円玉:2枚から合計金額がちょうど100円になる選び方は2通り",
			bytes.NewBufferString("2\n2\n2\n100"),
			2,
		},
		{
			"500円玉:5枚、100円玉:1枚、50円玉:0枚から合計金額がちょうど150円になる選び方は0通り",
			bytes.NewBufferString("5\n1\n0\n150"),
			0,
		},
		{
			"500円玉:30枚、100円玉:40枚、50円玉:50枚から合計金額がちょうど6000円になる選び方は213通り",
			bytes.NewBufferString("30\n40\n50\n6000"),
			213,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc087b(test.in); got != test.want {
				t.Errorf("abc087b(): got %v want %v", got, test.want)
			}
		})
	}

}
