package easy

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{0, 1, 1, 2, 2, 2, 3, 4, 5, 5, 6, 7, 7}

	num := RemoveDuplicates(input)

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

	output1 := MaxProfit(input1)
	output2 := MaxProfit(input2)
	output3 := MaxProfit(input3)

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
	input := []int{1, 2, 3, 4, 5, 6, 7}
	k := 2
	expected := []int{6, 7, 1, 2, 3, 4, 5}
	Rotate(input, k)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("expected the rusult of : %v to be %d but instead got %d!", []int{1, 2, 3, 4, 5, 6, 7}, expected, input)
	}
}

func TestContainsDuplicate(t *testing.T) {
	input1 := []int{1, 2, 3, 1}
	e1 := ContainsDuplicate(input1)

	input2 := []int{1, 2, 3, 4}
	e2 := ContainsDuplicate(input2)

	if !e1 {
		t.Errorf("expected the result of %v to be true but insteaded got %v!", input1, e1)
	}

	if e2 {
		t.Errorf("expected the result of %v to be false but insteaded got %v!", input2, e2)
	}
}

func TestSingleNumber(t *testing.T) {
	input := []int{1, 2, 3, 1, 2}
	output := SingleNumber(input)
	expected := 3
	if output != expected {
		t.Errorf("expected the result of %v to be %d but insteaded got %d!", input, expected, output)
	}
}

func TestPlusOne(t *testing.T) {
	input := []int{9,9}
	output := PlusOne(input)
	expected := []int{1,0,0}
	if ! reflect.DeepEqual(output, expected) {
		t.Errorf("Expected the result of %v to be %v but insteaded got %v!", input, expected, output)
	}
}

func TestMoveZeros(t *testing.T) {
	input := []int{0, 3, 0, 0, 0, 1, 2, 7, 0, 3, 1}
	MoveZeros(input)
	expected := []int{3,1,2,7,3,1,0,0,0,0,0}

	if ! reflect.DeepEqual(input, expected) {
		t.Errorf("Expected the result of %v to be %v but insteaded got %v!", []int{0, 3, 0, 0, 0, 1, 2, 7, 0, 3, 1}, expected, input)

	}
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	res := TwoSum(nums, target)

	if nums[res[0]] + nums[res[1]] != target {
		t.Errorf("Expect nums's (%v) nums[%d] + nums[%d] = %d but instead of got %d",nums, res[0], res[1],target, nums[res[0]] + nums[res[1]])
	}
}

func TestIsValidSudoku(t *testing.T) {
	/*
	[[".",".","5",".",".",".",".",".","."],
	["1",".",".","2",".",".",".",".","."],
	[".",".","6",".",".","3",".",".","."],
	["8",".",".",".",".",".",".",".","."],
	["3",".","1","5","2",".",".",".","."],
	[".",".",".",".",".",".",".","4","."],
	[".",".","6",".",".",".",".",".","."],
	[".",".",".",".",".",".",".","9","."],
	[".",".",".",".",".",".",".",".","."]]

	[
	[".",".","5",".",".",".",".",".","."],
	["1",".",".","2",".",".",".",".","."],
	[".",".","6",".",".","3",".",".","."],
	["8",".",".",".",".",".",".",".","."],
	["3",".","1","5","2",".",".",".","."],
	[".",".",".",".",".",".",".","4","."],
	[".",".","6",".",".",".",".",".","."],
	[".",".",".",".",".",".",".","9","."],
	[".",".",".",".",".",".",".",".","."]
	]
	 */
	input := [][]byte{
		{'.','.','5','.','.','.','.','.','.'},
		{'1','.','.','2','.','.','.','.','.'},
		{'.','.','6','.','.','3','.','.','.'},
		{'8','.','.','.','.','.','.','.','.'},
		{'3','.','1','5','2','.','.','.','.'},
		{'.','.','.','.','.','.','.','4','.'},
		{'.','.','6','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','9','.'},
		{'.','.','.','.','.','.','.','.','.'},
	}

	res := IsValidSudoku(input)

	if res {
		t.Errorf("Somethis wrong %v", input)
	}
}

func TestMatrixRotate(t *testing.T) {
	matrix := make([][]int, 3)
	for i:=0; i<3; i++ {
		matrix[i] = []int{i+1,i+2,i+3}
	}

	expected := [][]int{
		{3,2,1},
		{4,3,2},
		{5,4,3},
	}
	MatrixRotate(matrix)

	if !reflect.DeepEqual(matrix, expected) {
		t.Errorf("Something wrong and I don't want to say anything")
	}

}