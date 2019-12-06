package medium

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	//input := []int{1,2,3,4,5,6,7,8,9,9,12,19,19,20}
	input := []int{2,2}
	fmt.Println(SearchRange(input, 2))
}

func TestMerge(t *testing.T) {
	input := [][]int{
		{1,3},
		{2,6},
		{8,10},
		{15,18},
	}
	fmt.Println(Merge(input))
}

func TestSearchAndReverse(t *testing.T) {
	input := []int{3,1}
	fmt.Println(SearchAndReverse(input, 1))
}

func TestSearchMatrix(t *testing.T) {
	inputs := [][]int{
		{1,   4,  7, 11, 15},
		{2,   5,  8, 12, 19},
		{3,   6,  9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(SearchMatrix(inputs, 1))
	fmt.Println(SearchMatrix(inputs, 15))
	fmt.Println(SearchMatrix(inputs, 3))
	fmt.Println(SearchMatrix(inputs, 9))
	fmt.Println(SearchMatrix(inputs, 23))
	fmt.Println(SearchMatrix(inputs, 26))
	fmt.Println(SearchMatrix(inputs, 31))
	fmt.Println(SearchMatrix(inputs, 10))


}
