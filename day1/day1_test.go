package day1

import "testing"

func TestParseFileByNewline(t *testing.T) {
	input := "input_test.txt"
	expected := []int{159, 158, 174}
	actual := parseFile(input)
	if len(actual) != len(expected) {
		t.Error("Expected ", len(expected), " but got ", len(actual))
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Error("Expected ", expected[i], " but got ", actual[i])
		}
	}
}

func TestIsPrevLessThan(t *testing.T) {
	if !isPrevLessThan(1, 2) {
		t.Error("Expected true but got false")
	}
	if isPrevLessThan(2, 1) {
		t.Error("Expected false but got true")
	}
	if isPrevLessThan(1, 1) {
		t.Error("Expected false but got true")
	}
}

func TestGetIncreaseCount(t *testing.T) {
	increase := getIncreaseCount("input_test.txt")
	if increase != 3 {
		t.Error("Expected \"3!\" but got ", increase)
	}
}
