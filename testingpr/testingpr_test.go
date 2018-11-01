package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 == 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{3, 5},
		{4, 6},
		{5, 7},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test failed: {} inputted, {} expected, received {}", test.input, test.expected, output)
		}
	}
}

func TestTableSum(t *testing.T) {
	var sumTests = []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 2, 3},
		{3, 3, 6},
		{4, 5, 9},
		{1, 7, 8},
		{0, 7, 7},
	}

	for _, test := range sumTests {
		if output := Sum(test.num1, test.num2); output != test.expected {
			t.Error("Test failed: {}, {} inputted, {} expected, received {}", test.num1, test.num2, test.expected, output)
		}
	}
}
