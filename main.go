package main

import "fmt"

func main() {
	fmt.Println("My LeetCode story:")

	input := []int{0, 3, 0, 0, 0, 1, 2, 7, 0, 3, 1}
	moveZeros(input)
	fmt.Println(input)
}
