package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{
			data:   []int{21, 21},
			answer: 42},
		test{
			data:   []int{5, 10},
			answer: 15},
		test{
			data:   []int{12, 10},
			answer: 22},
		test{
			data:   []int{-4, 10},
			answer: 6},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
