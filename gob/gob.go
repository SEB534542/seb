package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"io/ioutil"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func AddToGobFile(i interface{}, fname string) error {
	// Initialize  encoder
	var data bytes.Buffer
	enc := gob.NewEncoder(&data) // Will encode (write) to data

	// Encode (send) some values.
	err := enc.Encode(i)
	if err != nil {
		return fmt.Errorf("Encode error: %v", err)
	}

	// Store encoded data in file fname
	err = ioutil.WriteFile(fname, data.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("Write error to '%v': %v", fname, err)
	}
	return nil
}

func ReadGob(i interface{}, fname string) {
	// Initialize decoder
	var data bytes.Buffer
	dec := gob.NewDecoder(&data) // Will decode (read) and store into data

	// Read content from file
	content, err := ioutil.ReadFile("test.gob")
	if err != nil {
		log.Fatal("Error reading file '%v':", fname, err)
	}
	y := bytes.NewBuffer(content)
	data = *y

	// Decode (receive) and print the values.

	err = dec.Decode(i)
	if err != nil {
		log.Fatal("Decode error:", err)
	}

}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func main() {

	AddToGobFile(P{3, 4, 5, "Pythagoras"}, "test2.gob")

	var q Q
	ReadGob(&q, "test2.gob")
	fmt.Println(q)

}
