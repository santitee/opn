package assignment4_test

import (
	"opn/assignment4"
	"testing"
)

func TestCalStudentGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "A", score: 90, expected: "A"},
		{name: "B", score: 80, expected: "B"},
		{name: "C", score: 70, expected: "C"},
		{name: "D", score: 60, expected: "D"},
		{name: "E", score: 50, expected: "E"},
		{name: "F", score: 48, expected: "F"},
		{name: "can not calculate please check score value", score: 101, expected: "can not calculate please check score value"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := assignment4.CalStudentGrade(c.score)

			if grade != c.expected {
				t.Errorf("got %v expected %v", grade, c.expected)
			}
		})
	}
}
