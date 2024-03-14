package Easy

func romanToInt(s string) int {
	totalSum := 0
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i, i2 := range s {
		totalSum += romanMap[i2]

		if i != 0 && romanMap[rune(s[i-1])] < romanMap[rune(s[i])] {
			totalSum -= 2 * romanMap[rune(s[i-1])]
		}
	}
	return totalSum
}
