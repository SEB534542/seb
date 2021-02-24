package main

import (
	"fmt"

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

	a := P{3, 4, 5, "Pythagoras"}

	// Save var a (which is of type P)
	seb.SaveGob(a, "test.gob")

	// load var a into struct P
	var p P
	seb.ReadGob(&p, "test.gob")
	fmt.Println("Loaded var p:", p.X, p.Y, p.Z, p.Name)
	fmt.Printf("Types: %T / %T\n", p, p.X)

	// load var a into struct Q
	var q Q
	seb.ReadGob(&q, "test.gob")
	fmt.Println("Loaded var q:", q.X, q.Y, p.Z, q.Name)
	fmt.Printf("Types: %T / %T\n", q, q.X)
}
