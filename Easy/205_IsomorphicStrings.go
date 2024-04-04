package Easy

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	map1 := make([]int, 128)
	map2 := make([]int, 128)

	for i := 0; i < len(s); i++ {
		sch := s[i]
		tch := t[i]

		if map1[sch] == 0 && map2[tch] == 0 {
			map1[sch] = int(tch)
			map2[tch] = int(sch)
		} else if map1[sch] != int(tch) || map2[tch] != int(sch) {
			return false
		}
	}
	return true
}
