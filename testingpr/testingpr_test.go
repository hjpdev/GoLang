package main

import "testing"

func TestPlus2(t *testing.T) {
	if Plus2(2) != 4 {
		t.Error("Expected 2 + 2 == 4")
	}
}

func TestTablePlus2(t *testing.T) {
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
		if output := Plus2(test.input); output != test.expected {
			t.Errorf("Test failed: %d inputted, %d expected, received %d", test.input, test.expected, output)
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
			t.Errorf("Test failed: %d, %d inputted, %d expected, received %d", test.num1, test.num2, test.expected, output)
		}
	}
}

func TestTableMinus(t *testing.T) {
	var minusTests = []struct {
		num1     int
		num2     int
		expected int
	}{
		{2, 1, 1},
		{7, 5, 2},
		{2, 2, 0},
		{10, 12, -2},
		{5, 4, 1},
		{2, 7, -5},
	}

	for _, test := range minusTests {
		if output := Minus(test.num1, test.num2); output != test.expected {
			t.Errorf("Test failed: %d, %d inputted, %d expected, received %d", test.num1, test.num2, test.expected, output)
		}
	}
}

func TestTableDivider(t *testing.T) {
	var divisionTests = []struct {
		num1     float32
		num2     float32
		expected float32
	}{
		{9.0, 3.0, 3.0},
		{12.0, 4.0, 3.0},
		{9.0, 9.0, 1.0},
		{10.0, 5.0, 2.0},
		{2.0, 1.0, 2.0},
	}

	for _, test := range divisionTests {
		if output := Divider(test.num1, test.num2); output != test.expected {
			t.Errorf("Test failed: %f, %f inputted, %f expected, receivec %f", test.num1, test.num2, test.expected, output)
		}
	}
}

func TestTablePower(t *testing.T) {
	var powerTests = []struct {
		num1     int
		num2     int
		expected int
	}{
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 8},
		{3, 3, 27},
		{1, 1, 1},
	}

	for _, test := range powerTests {
		if output := Power(test.num1, test.num2); output != test.expected {
			t.Errorf("Test failed: %d, %d inputted, %d expected, received %d", test.num1, test.num2, test.expected, output)
		}
	}
}
