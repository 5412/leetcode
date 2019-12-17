package medium

import (
	"fmt"
	"leetcode/structs"
	"testing"
)

func TestSerialize(t *testing.T) {
	input := []int{1,2,3,4,5}
	root := structs.GeneralBinaryTree(input)
	root2 := structs.TreeNode{Val:1, Left:nil, Right:nil}
	node1 := structs.TreeNode{Val:2, Left:nil, Right:nil}
	node2 := structs.TreeNode{Val:3, Left:nil, Right:nil}
	node4 := structs.TreeNode{Val:4, Left:nil, Right:nil}
	node5 := structs.TreeNode{Val:5, Left:nil, Right:nil}
	root2.Left = &node1
	root2.Right = &node2
	node2.Left = &node4
	node2.Right = &node5
	fmt.Println("!23123")
	fmt.Println(Serialize(root))
	str := "1,2,3,null,null,4,5,null,null,null,null"
	fmt.Println(Unserialize(str))
	root3 := Unserialize(str)
	fmt.Println(root3.Val)
	fmt.Println(root3.Left.Val)
	fmt.Println(root3.Right.Val)
	fmt.Println(root3.Left.Left)
	fmt.Println(root3.Left.Right)
	fmt.Println(root3.Left)
	fmt.Println(root3.Right)

}
