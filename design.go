package leetcode

//import "math/rand"
//
//type Solution struct {
//	origin []int
//	shuffled []int
//	len int
//}
//
//func Constructor(nums []int) Solution {
//	l := len(nums)
//	return Solution{origin:nums,len:l}
//}
//
//func (this *Solution) Reset() []int {
// 	return this.origin
//}
//
//func (this *Solution) Shuffle() []int {
//	if this.len < 2 {
//		return nil
//	}
//	index1 := rand.Int() % this.len
//	index2 := rand.Int() % this.len
//	this.shuffled[index1], this.shuffled[index2] = this.shuffled[index2], this.shuffled[index1]
//	return this.shuffled
//}

// God's algorithm

//type Solution struct {
//	array  []int
//	backup []int //备份
//}
//
//func Constructor(nums []int) Solution {
//	return Solution{
//		backup: append([]int{}, nums...),
//		array:  append([]int{}, nums...),
//	}
//}
//
//// Reset the array to its original configuration and return it.
//func (this *Solution) Reset() []int {
//	this.array = append([]int{},this.backup...)
//	return this.backup
//}
//
//
//func (this *Solution) Shuffle() []int {
//	for i := len(this.array); i>0; i-- {
//		p := rand.Intn(i)
//		this.array[i-1],this.array[p]=this.array[p],this.array[i-1]
//	}
//	return this.array
//}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

//type MinStack struct {
//	elements []int
//	min int
//	l int
//}
//func Constructor() MinStack {
//	return MinStack{min: math.MaxInt64}
//}
//
//
//func (this *MinStack) Push(x int)  {
//	this.elements = append(this.elements, x)
//	this.l++
//	if x < this.min {
//		this.min = x
//	}
//}
//
//
//func (this *MinStack) Pop()  {
//	if this.l > 0 {
//		e := this.elements[this.l-1]
//		this.elements = this.elements[:this.l-1]
//		this.l--
//		if e == this.min{
//			this.min = math.MaxInt64
//			for i:=0; i<this.l; i++ {
//				if this.min > this.elements[i] {
//					this.min = this.elements[i]
//				}
//			}
//		}
//	}
//}
//
//
//func (this *MinStack) Top() int {
//	if this.l > 0 {
//		return this.elements[this.l - 1]
//	}
//	return 0
//}
//
//
//func (this *MinStack) GetMin() int {
//	return this.min
//}


// leetcode God's algorithm
// 每个栈元素由 元素值与其对应的最小值组成

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{min: min, x: x})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].x

}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */