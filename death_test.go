package death

import (
	"testing"
)

// By generates a function that kills the program by calling t.Fatal.
func By(t *testing.T) func(err error) {
	return func(err error) {
		if err == nil {
			return
		}
		t.Fatal(err)
	}
}
