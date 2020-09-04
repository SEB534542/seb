package seb

import (
	"fmt"
	"testing"
)

func TestStrToIntZ(t *testing.T) {
	type test struct {
		s    string // Input for func StrToIntZ
		want int    // Wanted output
		err  error  // Wanted error
	}
	tests := []test{
		{"5", 5, nil},
		{"-1", 0, nil},
		{"0", 0, nil},
		{"1.0", 0, nil},
		{"1.5", 0, fmt.Errorf("StrToIntZ: unable to transform %s to an int", "Test")},
		{"Test", 0, fmt.Errorf("StrToIntZ: unable to transform %s to an int", "Test")},
	}
	for _, v := range tests {
		got, err := StrToIntZ(v.s)
		if got != v.want && err != v.err {
			t.Error("Want:", v.want, "Got:", got, "Error:", err)
		}
	}
}

func TestReverseXS(t *testing.T) {
	type test struct {
		xs   []string
		want []string
	}
	tests := []test{
		{[]string{"a", "b", "c", "d", "e"}, []string{"e", "d", "c", "b", "a"}},
	}
	for _, v := range tests {
		got := ReverseXS(v.xs)
		for i, _ := range got {
			if got[i] != v.want[i] {
				t.Error("Want:", v.want, "Got:", got)
			}
		}
	}
}
