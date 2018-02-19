/*
Package death provides a tiny set of functions that I find myself duplicating time and time again in my Go programs.
*/
package death

import (
	"fmt"
	"os"
	"testing"
)

// Code provides the exit code for Main.
var Code = 1

// Main exits the program by printing an error to stderr (if err is not nil), and exiting with Code.
func Main(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(Code)
}

// By generates a function that kills the program by calling t.Fatal.
func By(t *testing.T) func(err error) {
	return func(err error) {
		if err == nil {
			return
		}
		t.Fatal(err)
	}
}
