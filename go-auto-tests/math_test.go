package main

import "testing"

func TestMath(t *testing.T) {
	result := Math(13, 50)
	if result != 650 {
		t.Error("wrong result", result)
	}
}
