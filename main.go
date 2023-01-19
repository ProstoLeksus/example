package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	array := [5]int{5, 7, 9, 2, 2}
	target := 11
	fmt.Println(id)
	a := twoSum(array, target)
	fmt.Println(a)
}

func twoSum(nums [5]int, target int) []int {
	result := []int{2, 3, 4}
	return result
}
