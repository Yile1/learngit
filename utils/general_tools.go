package utils

import "strings"

func IsStringInArray(target string, str_array []string) bool {
	for _, element := range str_array{
		if target == element{
			return true
		}
	}
	return false
}

func IsUpdateVersionAvailable(updateVersionCode string, minUpdateVersionCode string, maxUpdateVersionCode string) bool {
	// compare rules
	return true
}

func SplitStringToList(str string)  []string{
	return strings.Split(str,",")
}