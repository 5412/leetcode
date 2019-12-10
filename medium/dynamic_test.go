package medium

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	//inputs := []int{3,2,1,0,4}
	inputs := []int{2,3,1,1,4}
	fmt.Println("1----------1")
	fmt.Println(CanJump(inputs))
	fmt.Println("1----------1")
}

func TestUniquePaths(t *testing.T) {
	fmt.Println("----------")
	fmt.Println(UniquePaths(23,12))
	fmt.Println("----------")
}
