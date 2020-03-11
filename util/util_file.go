package util

import "strings"

func GetCurrentPath(fullPath string) string {
	index1 := strings.LastIndex(fullPath, "/")
	//maybe windows env
	index2 := strings.LastIndex(fullPath, "\\")
	index := index1
	if index2 > index1 {
		index = index2
	}
	return fullPath[index+1:]
}
