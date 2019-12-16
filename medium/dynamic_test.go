package medium

import (
	"fmt"
	"testing"
)

//func TestCanJump(t *testing.T) {
//	//inputs := []int{3,2,1,0,4}
//	inputs := []int{2,3,1,1,4}
//	fmt.Println("1----------1")
//	fmt.Println(CanJump(inputs))
//	fmt.Println("1----------1")
//}
//
//func TestUniquePaths(t *testing.T) {
//	fmt.Println("----------")
//	fmt.Println(UniquePaths(23,12))
//	fmt.Println("----------")
//}
//
//func TestCoinChange(t *testing.T) {
//	fmt.Println("2-----------------------2")
//	fmt.Println(CoinChange([]int{4, 3}, 10))
//	fmt.Println("2-----------------------2")
//}

func TestLengthOfLIS(t *testing.T) {
	fmt.Println(LengthOfLIS([]int{1,3,6,7,9,4,10,5,6}))
}