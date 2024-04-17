package Easy

func MaxDepth(s string) int {
	answer := 0
	open := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			open++
			if open > answer {
				answer = open
			}
		case ')':
			open--
		}
	}
	return answer
}
