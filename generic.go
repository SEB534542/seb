// Package seb contains generic functions
package seb

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"reflect"
	"strconv"
	"time"
)

// MaxIntSlice takes a variadic parameter of integers and
// returns the highest integer.
func MaxIntSlice(xi ...int) int {
	var max int
	for i, v := range xi {
		if i == 0 || v > max {
			max = v
		}
	}
	return max
}

// SaveToJson takes an interface, stores it into the filename
// and returns an error (or nil).
func SaveToJSON(i interface{}, fileName string) error {
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

// SendMail sends an e-mail to one or more recipients.
// Example: SendMail([]string("recipient1@test.com", "recipient2@test.com"), "sender@test.com", "Subject", "Body", "12345", "smtp.gmail.com", "587"))
func SendMail(to []string, from, subj, body, password, domain, port string) error {
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
		return fmt.Errorf("SendMail: error when sending: %v", err)
	}
	return nil
}

// ReadCSV reads a CSV file and returns the output as a slice of
// slice of string, where the main slice represents the rows and the subsequent
// slice the column values.
func ReadCSV(file string) [][]string {
	// Read the file
	f, err := os.Open(file)
	if err != nil {
		f, err := os.Create(file)
		if err != nil {
			log.Panic("Unable to create csv", err)
		}
		f.Close()
		return [][]string{}
	}
	defer f.Close()
	r := csv.NewReader(f)
	lines, err := r.ReadAll()
	if err != nil {
		log.Panic(err)
	}
	return lines
}

// AppendCSV takes a CSV filename and slice of new lines and
// adds the later to the existing CSV file.
func AppendCSV(file string, newLines [][]string) error {

	// Get current data
	lines := ReadCSV(file)

	// Add new lines
	lines = append(lines, newLines...)

	// Write the file
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	if err = w.WriteAll(lines); err != nil {
		return err
	}
	return nil
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

// ReverseXs takes a slice of string and returns the exact same
// slice in the the opposite order. If a is 10 long, then a[0]
// will be a[9] etc.
func ReverseXs(a []string) []string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// ReverseXxs does exactly the same as ReverseXs but now for a slice of slice
// of string, where the first (main) slice is reversed.
// This can be used for example on the output of ReadCSV.
func ReverseXss(a [][]string) [][]string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

// CalcAverage takes a variadic parameter of integers and
// returns the average integer.
func CalcAverage(xi ...int) int {
	total := 0
	for _, v := range xi {
		total = total + v
	}
	return total / len(xi)
}

// StoTime receives a string of time (format hh:mm) and a day offset,
// and returns a type time with today's and the supplied hours and
// minutes + the offset in days.
func StoTime(t string, days int) (time.Time, error) {
	timeNow := time.Now()

	timeHour, err := strconv.Atoi(t[:2])
	if err != nil {
		return time.Time{}, err
	}

	timeMinute, err := strconv.Atoi(t[3:])
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()+days, int(timeHour), int(timeMinute), 0, 0, time.Local), nil
}

// loadConfig loads configuration from a given json file (including folder) and loads it into i interface.
func LoadConfig(fname string, i interface{}) error {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		log.Printf("File '%v' does not exist, creating blank", fname)
		SaveToJSON(i, fname)
	} else {
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			return fmt.Errorf("%s is corrupt. Please delete the file (%v)", fname, err)
		}
		err = json.Unmarshal(data, i)
		if err != nil {
			return fmt.Errorf("%s is corrupt. Please delete the file (%v)", fname, err)
		}
	}
	return nil
}

// ReadGob reads a gob from a file and converts it into an interface.
func ReadGob(i interface{}, fname string) error {
	// Initialize decoder
	var data bytes.Buffer
	dec := gob.NewDecoder(&data) // Will decode (read) and store into data

	// Read content from file
	content, err := ioutil.ReadFile("test.gob")
	if err != nil {
		return fmt.Errorf("Error reading file '%v': %v", fname, err)
	}
	y := bytes.NewBuffer(content)
	data = *y

	// Decode (receive) and print the values.

	err = dec.Decode(i)
	if err != nil {
		return fmt.Errorf("Error decoding into '%v': %v (%v)", fname, err, i)
	}
	return nil
}

// SaveGob encodes an interface and stores it as a Gob into a file named fname.
func SaveToGob(i interface{}, fname string) error {
	var data bytes.Buffer

	enc := gob.NewEncoder(&data) // Will write to data
	//	dec := gob.NewDecoder(&data) // Will read from data

	// Encode (send) some values.
	err := enc.Encode(i)
	if err != nil {
		return fmt.Errorf("Error encoding '%v': %v", fname, err)
	}

	// Store data
	err = ioutil.WriteFile("test.gob", data.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("Error storing '%v': %v", fname, err)
	}
	return nil
}

// XlsxColNames takes a struct and returns a map of the xlsx tags in that struct (i.e. the columns) and the name of the associated struct field.
// It can be used to validate the imported column types from "github.com/tealeg/xlsx/". For example by using below code:
//
//  var nilConverter xlsx.CellVisitorFunc = func(c *xlsx.Cell) error {
// 	col, _ := c.GetCoordinates()
// 	string, _ := c.FormattedValue()
// 	if string == "" {
// 		switch colNames[col] {
// 		case "Days":
// 			c.SetInt(0)
// 		case "TotalDays":
// 			c.SetInt(0)
// 		case "FinEnd":
// 			c.SetDate(time.Date(9999, time.December, 31, 0, 0, 0, 0, time.UTC))
// 		case "NLOK1":
// 			c.SetFloat(0.0)
// 		case "NLOK2":
// 			c.SetFloat(0.0)
// 		case "NLPAY":
// 			c.SetFloat(0.0)
// 		case "NLPAYFTE":
// 			c.SetFloat(0.0)
// 		}
// 	}
// 	return nil
// 	}
//  file, err := xlsx.OpenFile(fname)
// 	if err != nil {
// 		return sr, fmt.Errorf("Error while opening file '%v': %v", fname, err)
// 	}
// 	for i := startRow; i <= file.Sheets[0].MaxRow-1; i++ {
// 		line := &sickness{}
// 		row, err := file.Sheets[0].Row(i)
// 		if err != nil {
// 			return sr, fmt.Errorf("Error retrieving row %v from '%v': %v", i+1, fname, err)
// 		}
// 		row.ForEachCell(nilConverter)
// 		err = row.ReadStruct(line)
// 		if err != nil {
// 			return sr, fmt.Errorf("Error converting row %v from '%v': %v", i+1, fname, err)
// 		}
// 		sr = append(sr, *line)
// 	}
func XlsxColNames(s interface{}) map[int]string {
	const tagName = "xlsx"
	m := map[int]string{}
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		tag := field.Tag.Get(tagName)
		if tag != "-" && tag != "" {
			col, err := strconv.Atoi(tag)
			if err == nil {
				m[col] = field.Name
			}
		}
	}
	return m
}
