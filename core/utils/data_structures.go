package utils

import (
	"math/rand"
	"strings"
	"time"
)

func DeleteEmptyElements(source []string) []string {
	var result []string
	for _, str := range source {
		if strings.Trim(str, " ") != "" {
			result = append(result, str)
		}
	}
	return result
}

func GetRandomNumber(min, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min+1) + min)
}
