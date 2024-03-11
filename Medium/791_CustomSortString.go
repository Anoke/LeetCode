package main

import "strings"

func customSortString(order string, s string) string {
	mapa := make(map[int32]int)
	for _, i2 := range s {
		mapa[i2]++
	}

	result := ""
	for _, i2 := range order {
		if _, ok := mapa[i2]; ok && mapa[i2] != 0 {
			result = result + strings.Repeat(string(i2), mapa[i2])
			mapa[i2] = 0
		}
	}
	for _, i := range s {
		if mapa[i] != 0 {
			result = result + strings.Repeat(string(i), mapa[i])
			mapa[i] = 0
		}
	}
	return result
}
