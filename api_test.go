package main

import (
	"testing"
)

func TestGetFlag(t *testing.T) {
	loadConfigGlobals()

	flag, foundFlag, err := getFlag("test1", "flag that doesn't exist")
	if err != nil {
		t.Error("Failed to get a response: ", err)
	}
	if flag != false || foundFlag != false {
		t.Error("Failed to get a false value for a missing flag.")
	}
}
