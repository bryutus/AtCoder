package main

import (
	"bytes"
	"testing"
)

func TestAbc174a(t *testing.T) {
	testCases := []struct {
		desc string
		in   *bytes.Buffer
		want string
	}{
		{
			"室温が25度なので冷房の電源を入れない",
			bytes.NewBufferString("25"),
			"No",
		},
		{
			"室温が30度なので冷房の電源を入れる",
			bytes.NewBufferString("30"),
			"Yes",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			if got := abc174a(test.in); got != test.want {
				t.Errorf("abc174a(): got %v want %v", got, test.want)
			}
		})
	}
}
