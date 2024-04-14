package utils

import (
	"strconv"
	"strings"
)

func ConvertTagIDs(str string) []int {
	s := strings.Trim(str, "{}")
	parts := strings.Split(s, ",")
	tags := make([]int, len(parts))
	for i, part := range parts {
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil
		}
		tags[i] = n
	}
	return tags
}
