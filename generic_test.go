package seb

import (
	"fmt"
	"os"
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

func TestMaxIntSlice(t *testing.T) {
	type test struct {
		xi   []int
		want int
	}
	tests := []test{
		{[]int{0, 1, 2, 3, 4, 5, 6}, 6},
		{[]int{0, -1, -2, -3, -4, -5, -6}, 0},
		{[]int{0, -1, -2, 200, -4, -5, -6}, 200},
	}
	for _, v := range tests {
		got := MaxIntSlice(v.xi...)
		if v.want != got {
			t.Error("Want:", v.want, "Got:", got)
		}
	}
}

func TestAppendAndReadCSV(t *testing.T) {
	fname := "test.csv"
	os.Remove(fname)
	xxs1 := [][]string{
		[]string{"A", "B", "C"},
		[]string{"D", "E", "F"},
	}
	err := AppendCSV(fname, xxs1)
	if err != nil {
		t.Error("Error appending to file:", err)
	}
	xxs2 := ReadCSV(fname)
	for i, v := range xxs2 {
		for j, _ := range v {
			if xxs1[i][j] != xxs2[i][j] {
				t.Error("Error on line", i, "Want:", xxs1, "Got:", xxs2)
			}
		}
	}
	xxsNew := [][]string{
		[]string{"G", "H", "I"},
		[]string{"J", "K", "L"},
	}
	err = AppendCSV(fname, xxsNew)
	if err != nil {
		t.Error("Error appending to file:", err)
	}
	xxs3 := append(xxs1, xxsNew...)
	xxs4 := ReadCSV(fname)
	for i, v := range xxs4 {
		for j, _ := range v {
			if xxs4[i][j] != xxs3[i][j] {
				t.Error("Error on line", i, "Want:", xxs3, "Got:", xxs4)
			}
		}
	}
	os.Remove(fname)
}
