package Easy

//
//func isPalindrome(x int) bool {
//	str := strconv.Itoa(x)
//	if len(str) <= 1 {
//		return true
//	}
//	for i, _ := range str {
//		if i >= len(str)/2 {
//			return true
//		}
//		if str[i] != str[len(str)-1-i] {
//			return false
//		}
//	}
//	return true
//}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	temp := x
	for temp != 0 {
		reversed = reversed*10 + temp%10
		temp = temp / 10
	}
	if reversed == x {
		return true
	}
	return false
}
