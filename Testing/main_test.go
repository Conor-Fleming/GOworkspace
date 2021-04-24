package main

import "testing"

func TestMySum(t *testing.T) {
	if mySum(2, 3) != 5 {
		t.Error("Expected", 5, "Got", mySum(2, 3))
	}
}
