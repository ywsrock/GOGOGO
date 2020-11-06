package main

import "testing"

func TestIsBool(t *testing.T) {
	n := 1
	b := isBool(n)
	if b != true {
		t.Errorf("error ---> %d", n)
	}
}
