package Easy

func maxFrequencyElements(nums []int) int {
	dictionary := make(map[int]int)

	for _, num := range nums {
		dictionary[num]++
	}

	total := 0
	maxValue := 0

	for _, i2 := range dictionary {
		if i2 == maxValue {
			total += i2
		}
		if i2 > maxValue {
			maxValue = i2
			total = i2
		}
	}

	return total
}
