package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"

	"github.com/SEB534542/seb"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// This example shows how to store to and read gob from a file
func main() {

	SaveToGob(P{3, 4, 5, "Pythagoras"}, "test2.gob")

	var q Q
	seb.ReadGob(&q, "test2.gob")
	fmt.Println("Q:", q.X, q.Y, q.Name)
	fmt.Printf("Types: %T / %T\n", q, q.X)

	var p P
	seb.ReadGob(&p, "test2.gob")
	fmt.Println("P:", p.X, p.Y, p.Name)
	fmt.Printf("Types: %T / %T\n", p, p.X)
}

// SaveToGob encodes an interface and stores it into a file.
func SaveToGob(i interface{}, fname string) error {
	// Initialize  encoder
	var data bytes.Buffer
	enc := gob.NewEncoder(&data) // Will encode (write) to data

	// Encode (send) some values.
	fmt.Println(i)
	err := enc.Encode(i)
	if err != nil {
		return fmt.Errorf("Encode error: %v", err)
	}
	fmt.Println(data)
	fmt.Println(data.Bytes())

	// Store encoded data in file fname
	err = ioutil.WriteFile(fname, data.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("Write error to '%v': %v", fname, err)
	}
	return nil
}
