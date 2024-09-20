package utils

import (
	"strconv"
)

func Order(list []int) []int {
	newList := list

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if j+1 < len(newList) {
				current := newList[j]
				next := newList[j+1]
				if current > next {
					newList[j] = next
					newList[j+1] = current
				}
			}
		}
	}

	return newList
}

func Str2Int(str []string) ([]int, error) {
	result := []int{}
	for _, v := range str {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}

	return result, nil
}
