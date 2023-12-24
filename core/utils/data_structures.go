package utils

import "strings"

func DeleteEmptyElements(source []string) []string {
	var result []string
	for _, str := range source {
		if strings.Trim(str, " ") != "" {
			result = append(result, str)
		}
	}
	return result
}
