package Easy

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || char != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
