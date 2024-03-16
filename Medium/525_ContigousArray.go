package main

import "fmt"

func findMaxLength(nums []int) int {
	mapa := make(map[int]int)
	mapa[0] = -1
	count := 0
	maxlen := 0
	for i, num := range nums {
		if num == 1 {
			count++
		} else {
			count--
		}
		if _, ok := mapa[count]; ok {
			maxlen = max(maxlen, i-mapa[count])
		} else {
			mapa[count] = i
		}
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
}
