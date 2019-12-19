package medium

import (
	"leetcode/structs"
	"math/rand"
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
/**
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        if not root:
            return '#'
        return str(root.val)+','+self.serialize(root.left)+','+self.serialize(root.right)
    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        list = data.split(',')
        return self.deserializeTree(list)
    def deserializeTree(self, list):
        if len(list)<=0:
            return None
        val = list.pop(0)
        root = None
        if val != '#':
            root = TreeNode(int(val))
            root.left = self.deserializeTree(list)
            root.right = self.deserializeTree(list)
        return root

————————————————
版权声明：本文为CSDN博主「秋曾万」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_41679411/article/details/86503460
 */
func Serialize(root *structs.TreeNode) string {
	if root == nil {
		return "#"
	}
	val := strconv.Itoa(root.Val)
	return val + "," +  Serialize(root.Left) + "," + Serialize(root.Right)

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

/**
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
 */

/**
设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。

insert(val)：当元素 val 不存在时，向集合中插入该项。
remove(val)：元素 val 存在时，从集合中移除该项。
getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
示例 :

// 初始化一个空的集合。
RandomizedSet randomSet = new RandomizedSet();

// 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomSet.insert(1);

// 返回 false ，表示集合中不存在 2 。
randomSet.remove(2);

// 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomSet.insert(2);

// getRandom 应随机返回 1 或 2 。
randomSet.getRandom();

// 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomSet.remove(1);

// 2 已在集合中，所以返回 false 。
randomSet.insert(2);

// 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
randomSet.getRandom();
 */

//type RandomizedSet struct {
//	buf []int
//	m map[int]int
//}
//
//
///** Initialize your data structure here. */
//func Constructor() RandomizedSet {
//	return RandomizedSet{m:make(map[int]int),buf:make([]int,0)}
//}
//
//
///** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
//func (this *RandomizedSet) Insert(val int) bool {
//	if _, ok := this.m[val]; !ok {
//		this.buf = append(this.buf, val)
//		this.m[val] = len(this.buf) - 1
//		return true
//	}
//	return false
//}
//
//
///** Removes a value from the set. Returns true if the set contained the specified element. */
//func (this *RandomizedSet) Remove(val int) bool {
//	if index, ok := this.m[val]; ok {
//		l := len(this.buf)
//		this.m[this.buf[l-1]] = index
//		this.buf[index], this.buf[l-1] = this.buf[l-1], this.buf[index]
//		delete(this.m, val)
//		this.buf = this.buf[:l-1]
//		return true
//	}
//	return false
//}
//
//
///** Get a random element from the set. */
//func (this *RandomizedSet) GetRandom() int {
//	l := len(this.m)
//	random := rand.Int()
//	index := random % l
//	return this.buf[index]
//}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

/**
leetcode's algorithm
 */

type RandomizedSet struct {
	numsMap   map[int]int
	numsArray []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{numsArray: []int{}, numsMap: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numsMap[val]; ok {
		return false
	}

	this.numsMap[val] = len(this.numsArray)
	this.numsArray = append(this.numsArray, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.numsMap[val];
	if !ok {
		return false;
	}

	last := this.numsArray[len(this.numsArray)-1]
	if val == last { // if we are removing the last element in the array, we don't need to swap and update index in map
		this.numsArray = this.numsArray[:len(this.numsArray)-1]
	} else {
		// swap and remove the last element makes the delete of the element O(1)
		this.numsArray[index] = last
		this.numsArray = this.numsArray[:len(this.numsArray)-1]
		// note: we also need to update the index of the last element in the map
		this.numsMap[this.numsArray[index]] = index
	}
	delete(this.numsMap, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.numsArray))

	return this.numsArray[index]
}