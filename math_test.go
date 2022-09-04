package main

import "testing"

func TestSoma(t *testing.T) {
	some := Soma(10, 10)
	if some != 20 {
		t.Error("Soma should be 20")
	}
}
