package seb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func TestSaveAndReadJSON(t *testing.T) {
	const fname = "test.json"
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	test := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	SaveToJSON(test, fname)

	data, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Errorf("Error while reading %v: %v", fname, err)
	}

	want := ColorGroup{}

	err = json.Unmarshal(data, &want)
	if err != nil {
		t.Errorf("Error while unmarshalling %v: %v", fname, err)
	}

	if want.ID != test.ID || want.Name != test.Name {
		t.Error("Want:", want, "Got:", test)
	}

	for i, _ := range want.Colors {
		if want.Colors[i] != test.Colors[i] {
			t.Error("Error in colors.", "Want:", want, "Got:", test)
		}
	}
	os.Remove(fname)
}

func TestCalcAverage(t *testing.T) {
	type test struct {
		xi   []int
		want int
	}
	tests := []test{
		{[]int{0, 1, 2, 3, 4, 5, 6}, 3},
		{[]int{0, 8, 2, 100, 4, -4, 12}, 17}, // Average is float of 17.42857, which is int of 17
		{[]int{0, 8, 2, 100, 4, -4, 14}, 17}, // Average is float of 17.71429, which is int of 17
		{[]int{0}, 0},
		{[]int{-12, 8, 5, 4}, 1}, // Average is float of 1,25, which is int of 1
	}

	for _, v := range tests {
		got := CalcAverage(v.xi...)
		if got != v.want {
			t.Error("Want:", v.want, "Got:", got)
		}
	}

}
