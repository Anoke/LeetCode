package main

func productExceptSelf(nums []int) []int {
	total := 1
	zero_count := 0
	for _, num := range nums {
		if num != 0 {
			total *= num
		} else {
			zero_count++
		}
	}
	var result []int

	if zero_count > 1 {
		return make([]int, len(nums))
	}

	for _, num := range nums {
		if num != 0 {
			if zero_count == 1 {
				result = append(result, 0)
			} else {
				result = append(result, total/num)
			}
		} else {
			result = append(result, total)
		}
	}
	return result
}
