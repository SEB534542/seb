// Package seb contains generic functions
package seb

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strconv"
)

// MaxIntSlice takes a variadic parameter of integers and
// returns the highest integer
func MaxIntSlice(xi ...int) int {
	var max int
	for i, v := range xi {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

// SaveToJson takes an interface, stores it into the filename and returns an error
func SaveToJson(i interface{}, fileName string) error {
	bs, err := json.Marshal(i)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, bs, 0644)
	if err != nil {
		return err
	}
	return nil
}

// SendMail sends an e-mail to one or more recipients. Example:
// SendMail([]string("recipient1@test.com", "recipient2@test.com"), "sender@test.com", "Subject", "Body", "12345", "smtp.gmail.com", "587")
func SendMail(to []string, from, subj, body, password, domain, port string) {
	var msgTo string
	for i, s := range to {
		if i != 0 {
			msgTo = msgTo + ","
		}
		msgTo = msgTo + s
	}

	msg := []byte("To:" + msgTo + "\r\n" +
		"Subject:" + subj + "\r\n" +
		"\r\n" + body + "\r\n")

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, domain)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(domain+password, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadCSV reads a CSV file and returns the output as a slice of
// string.
func ReadCSV(file string) [][]string {
	// Read the file
	f, err := os.Open(file)
	if err != nil {
		f, err := os.Create(file)
		if err != nil {
			log.Fatal("Unable to create csv", err)
		}
		f.Close()
		return [][]string{}
	}
	defer f.Close()
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return lines
}

// AppendCSV takes a CSV filename and slice of new lines and
// adds the later to the existing CSV file.
func AppendCSV(file string, newLines [][]string) {

	// Get current data
	lines := ReadCSV(file)

	// Add new lines
	lines = append(lines, newLines...)

	// Write the file
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		log.Fatal(err)
	}
}

// StrToInt transforms string to an int and
// returns a positive int or zero.
func StrToIntZ(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("StrToIntZ: unable to transform %s to an int", s)
	}
	if i < 0 {
		return 0, err
	}
	return i, err
}

// Reverse XS takes a slice of string and returns the exact same
// slice in the the opposite order. If xs is 10 long, then xs[0]
// will be xs[9] etc.
func ReverseXS(xs []string) []string {
	r := []string{}
	for i, _ := range xs {
		r = append(r, xs[len(xs)-1-i])
	}
	return r
}

// ReverseXXS does exactly the same as ReverseXS but now for a slice of slice
// of string. This can be used for example on the output of ReadCSV.
func ReverseXSS(xxs [][]string) [][]string {
	r := [][]string{}
	for i, _ := range xxs {
		r = append(r, xxs[len(xxs)-1-i])
	}
	return r
}

// CalcAverage takes a variadic parameter of integers and
// returns the average integer
func CalcAverage(xi ...int) int {
	total := 0
	for _, v := range xi {
		total = total + v
	}
	return total / len(xi)
}
