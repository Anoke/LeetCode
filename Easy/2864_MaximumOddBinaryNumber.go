package Easy

import "strings"

func maximumOddBinaryNumber(s string) string {
	quantityOne := strings.Count(s, "1")
	indexLast := strings.LastIndex(s, "1")
	if quantityOne == 1 {
		s = s[:indexLast] + s[indexLast+1:] + "1"
		return s
	}
	one := "1"
	zero := "0"
	return strings.Repeat(one, quantityOne-1) + strings.Repeat(zero, len(s)-quantityOne) + "1"
}
