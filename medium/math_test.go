package medium

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	fmt.Println(IsHappy(19))
}

func TestTrailingZeroes(t *testing.T) {
	fmt.Println(TrailingZeroes(3))
	fmt.Println(TrailingZeroes(5))
	fmt.Println(TrailingZeroes(6))
	fmt.Println(TrailingZeroes(8))
	fmt.Println(TrailingZeroes(10))
	fmt.Println(TrailingZeroes(30))
}
