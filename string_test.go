package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	input := "hello"
	expected := "olleh"
	output := ReverseString(input)

	if output != expected {
		t.Errorf("Expected reverse of string %v is %v instead of got %v", input, expected, output)
	}
}

func TestReverseInt(t *testing.T) {
	input := 123
	expected := 321
	output := ReverseInt(input)

	if expected != output {
		t.Errorf("Expected reverse of int %v is %v,but instead of got %v!", input, expected, output)
	}
}
