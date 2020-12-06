package main

import (
	"fmt"
)

func main() {
	fmt.Print(twoSum([]int{1, 2, 3, 4}, 6))
}

func twoSum(nums []int, target int) []int {
	mapIndex := make(map[int]int)
	result := make([]int, 0, 2)

	for i, v := range nums {
		mapIndex[v] = i
	}
	for i := 0; i < len(nums)-1; i++ {
		currDiff := target - int(nums[i])

		if _, ok := mapIndex[currDiff]; ok != false && mapIndex[currDiff] != i {
			result = append(result, i, mapIndex[currDiff])
			return result
		}
	}
	return result
}

fmt.Print()