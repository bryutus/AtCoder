package main

import (
	"bytes"
	"testing"
)

func TestAbc171a(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want string
	}{
		{
			"Bは英大文字であるため、Aと出力される",
			bytes.NewBufferString(`B`),
			"A",
		},
		{
			"aは英大文字であるため、aと出力される",
			bytes.NewBufferString(`a`),
			"a",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc171a(test.in); got != test.want {
				t.Errorf("abc171a(): got %v want %v", got, test.want)
			}
		})
	}
}
