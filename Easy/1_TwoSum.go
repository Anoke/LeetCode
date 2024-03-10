package Easy

func twoSum(nums []int, target int) []int {
	mapa := make(map[int]int)
	for i, num := range nums {
		x := target - num
		if value, ok := mapa[x]; ok {
			return []int{i, value}
		}
		mapa[i] = num
	}
	return nil
}
