package main

import (
	"testing"
	"reflect"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 1, 1, 2, 2, 2, 3, 4, 5, 5, 6, 7, 7}

	num := removeDuplicates(input)

	expected := 8

	if num != expected {
		t.Errorf("expected the deplicate element num of : %v to be %d but instead got %d!", input, expected, num)
	}
}

func TestMaxProfit(t *testing.T) {
	input1 := []int{7, 6, 4, 3, 1}
	input2 := []int{1, 2, 3, 4, 5}
	input3 := []int{7, 1, 5, 3, 6, 4}

	expected1 := 0
	expected2 := 4
	expected3 := 7

	output1 := maxProfit(input1)
	output2 := maxProfit(input2)
	output3 := maxProfit(input3)

	switch {
	case expected1 != output1:
		t.Errorf("expected the profit of : %v to be %d but instead got %d!", input1, expected1, output1)
		fallthrough
	case expected2 != output2:
		t.Errorf("expected the profit of : %v to be %d but instead got %d!", input2, expected2, output2)
		fallthrough
	case expected3 != output3:
		t.Errorf("expected the profit of : %v to be %d but instead got %d!", input3, expected3, output3)
	}
}

func TestRotate(t *testing.T) {
	input := []int{1,2,3,4,5,6,7}
	k := 2
	expected := []int{6,7,1,2,3,4,5}
	rotate(input, k)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("expected the rusult of : %v to be %d but instead got %d!", []int{1,2,3,4,5,6,7}, expected, input)
	}
}

func TestContainsDuplicate(t *testing.T)  {
	input1 := []int{1,2,3,1}
	e1 := containsDuplicate(input1)

	input2 := []int{1,2,3,4}
	e2 := containsDuplicate(input2)

	if !e1 {
		t.Errorf("expected the result of %v to be true but insteaded got %v!", input1, e1)
	}

	if e2 {
		t.Errorf("expected the result of %v to be false but insteaded got %v!", input2, e2)
	}


}