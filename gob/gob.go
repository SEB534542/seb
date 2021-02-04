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

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func main() {
	// Initialize the encoder and decoder. Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var data bytes.Buffer

	enc := gob.NewEncoder(&data) // Will write to data
	dec := gob.NewDecoder(&data) // Will read from data

	// Encode (send) some values.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// Store data
	err = ioutil.WriteFile("test.gob", data.Bytes(), 0644)
	if err != nil {
		log.Fatal("error storing:", err)
	}

	// Read data
	raw, err := ioutil.ReadFile("test.gob")
	if err != nil {
		log.Fatal("error loading:", err)
	}
	x := &raw
	fmt.Printf("%T\n", *x)
	y := *x
	data = bytes.NewBuffer(y)

	// Decode (receive) and print the values.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

}
