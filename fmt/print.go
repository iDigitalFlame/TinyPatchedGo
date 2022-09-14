// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// 2022 - Shrinkage by iDigitalFlame

package fmt

import (
	"errors"
	"io"
)

// State represents the printer state passed to custom formatters.
// It provides access to the io.Writer interface plus information about
// the flags and options for the operand's format specifier.
type State interface {
	Write([]byte) (int, error)
	Width() (int, bool)
	Precision() (int, bool)
	Flag(int) bool
}

// Stringer is implemented by any value that has a String method,
// which defines the ``native'' format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an un-formatted printer
// such as Print.
type Stringer interface {
	String() string
}

// Formatter is implemented by any value thaa has a Format method.
// The implementation controls how State and rune are interpreted,
// and may call Sprint(f) or Fprint(f) etc. to generate its output.
type Formatter interface {
	Format(State, rune)
}

// GoStringer is implemented by any value that has a GoString method,
// which defines the Go syntax for that value.
// The GoString method is used to print values passed as an operand
// to a %#v format.
type GoStringer interface {
	GoString() string
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(v ...interface{}) string {
	return quickPrint(false, v...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintln(v ...interface{}) string {
	return quickPrint(true, v...)
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(v ...interface{}) (int, error) {
	return 0, nil
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(v ...interface{}) (int, error) {
	return 0, nil
}

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error.
//
// If the format specifier includes a %w verb with an error operand,
// the returned error will implement an Unwrap method returning the operand. It is
// invalid to include more than one %w verb or to supply it with an operand
// that does not implement the error interface. The %w verb is otherwise
// a synonym for %v.
func Errorf(s string, v ...interface{}) error {
	return errors.New(quickPrintf(s, v...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(s string, v ...interface{}) string {
	return quickPrintf(s, v...)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(s string, v ...interface{}) (int, error) {
	return 0, nil
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(w io.Writer, v ...interface{}) (int, error) {
	return quickFprint(w, false, v...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Fprintln(w io.Writer, v ...interface{}) (int, error) {
	return quickFprint(w, true, v...)
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func Fprintf(w io.Writer, s string, v ...interface{}) (int, error) {
	return quickFprintf(w, s, v...)
}
