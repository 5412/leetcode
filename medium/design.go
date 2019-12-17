package medium

import (
	"leetcode/structs"
	"strconv"
	"strings"
)

/**
序列化是将一个数据结构或者对象转换为连续的比特位的操作，
进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，
采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。
		这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:
你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
*/

func Serialize(root *structs.TreeNode) string {
	buf := make([]*structs.TreeNode, 0)
	buf = append(buf, root)
	res := make([]string, 0)
	for len(buf) != 0 {
		node := buf[0]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			buf = append(buf, node.Left)
			buf = append(buf, node.Right)
		} else {
			res = append(res, "null")
		}
		buf = buf[1:]
	}
	return strings.Join(res, ",")
}

func Unserialize(data string) *structs.TreeNode {
	var (
		nodes []*structs.TreeNode
		root, node  *structs.TreeNode
	)
	temp := strings.Split(data, ",")
	for i:=0; i<len(temp); i++ {
		if temp[i] != "null" {
			val, _ := strconv.Atoi(temp[i])
			node = &structs.TreeNode{Val:val, Left:nil, Right:nil}
		} else {
			node = nil
		}
		if root == nil {
			root = node
		} else {
			if len(nodes) == 0 {
				break
			}
			if nodes[0] != nil {
				if i%2 == 1 {
					nodes[0].Left = node
				} else {
					nodes[0].Right = node
					nodes = nodes[1:]
				}
			} else {
				nodes = nodes[1:]
			}
		}
		nodes = append(nodes, node)
	}
	return root
}
