package main

import (
	"bytes"
	"testing"
)

func TestAbc049c(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want string
	}{
		{
			"",
			bytes.NewBufferString("erasedream"),
			"YES",
		},
		{
			"",
			bytes.NewBufferString("dreameraser"),
			"YES",
		},
		{
			"",
			bytes.NewBufferString("dreamerer"),
			"NO",
		},
		{
			"",
			bytes.NewBufferString("dreraseeam"),
			"NO",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc049c(test.in); got != test.want {
				t.Errorf("abc049c(): got %v want %v", got, test.want)
			}
		})
	}
}
