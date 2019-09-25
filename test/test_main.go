package main

import "testing"
import "../src/test.go"

func TestExec(t *testing.T) {
	got := execConfig("")
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
