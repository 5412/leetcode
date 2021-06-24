package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCanCross(t *testing.T) {
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}

func TestIsCousins(t *testing.T) {
	root := TreeNode{
		Val: 1,
	}
	root.Left = nil
	node1 := TreeNode{Val: 2}
	root.Right = &node1
	node2 := TreeNode{Val: 3}
	node1.Right = &node2
	node3 := TreeNode{Val: 4}
	node2.Right = &node3
	node4 := TreeNode{Val: 5}
	node3.Right = &node4
	fmt.Println(isCousins(&root, 1, 3))
}

func TestStrangePrinter(t *testing.T) {
	input := "aaabbbcccsdbsadsqrrrrssbq"
	fmt.Println(strangePrinter(input))
}

func TestReverseParentheses(t *testing.T) {
	input := "(abcd)"
	fmt.Println(reverseParentheses(input))
}

func TestSearch(t *testing.T) {
	input := []int{1, 3, 5}
	fmt.Println(search(input, 1))
}

func TestPermutation(t *testing.T) {
	input := "abc"
	expect := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	got := permutation(input)
	if !reflect.DeepEqual(expect, got) {
		t.Fatalf("permutation err, expect=%v, go=%v", expect, got)
	}
}
