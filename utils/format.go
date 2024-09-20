package utils

import (
	"strings"
)

func Format(data []byte) ([]string) {
	str := string(data)
	str = strings.ReplaceAll(str, "\r\n", "\n")
	result := strings.Split(str, "\n")
	return result
}
