package main

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

//
//type Person struct {
//	Name  string
//	Phone string
//}
//
//func sleep(seconds int) {
//	<-time.After(time.Duration(seconds) * time.Second)
//}
//func main() {
//	var a, b int32 = math.MaxInt32, 1
//	sum := a + b
//	if (a > 0 && b > 0 && sum < 0) || (a < 0 && b < 0 && sum > 0) {
//		fmt.Println("Произошло переполнение!")
//	} else {
//		fmt.Println("Результат сложения:", sum)
//	}
//}
//
//// Функция для безопасного сложения двух int32 с проверкой на переполнение
//func safeAdd(a, b int32) (int32, bool) {
//	sum := a + b
//	if (a > 0 && b > 0 && sum < 0) || (a < 0 && b < 0 && sum > 0) {
//		fmt.Println("Произошло переполнение!")
//	} else {
//		fmt.Println("Результат сложения:", sum)
//	}
//}
