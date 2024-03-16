package main

import (
	"LeetCode/Easy"
	"fmt"
)

func numSubarraysWithSum(nums []int, goal int) int {
	totalCount := 0
	currentSum := 0
	freq := make(map[int]int)
	for _, num := range nums {
		currentSum += num
		if currentSum == goal {
			totalCount++
		}
		if _, ok := freq[currentSum-goal]; ok {
			totalCount += freq[currentSum-goal]
		}
		freq[currentSum]++
	}
	return totalCount
}

func main() {
	nums := []int{1, 0, 1, 0, 1}
	fmt.Println(numSubarraysWithSum(nums, 2))
	fmt.Println(Easy.RomanToInt("VII"))
}
