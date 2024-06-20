package utils

import (
	"os"
	"strings"
)

/*
description: gets the value of an environment variable
arguments:

	environVarName: the name of the variable to get the value of

return:
  - the value of the variable
  - an empty string if there is no such variable
*/
func GetEnvironVar(environVarName string) string {
	environVars := os.Environ()

	for _, environVar := range environVars {
		name := strings.Split(environVar, "=")[0]
		value := strings.Split(environVar, "=")[1]

		if name == environVarName {
			return value
		}
	}

	return ""
}

/*
description: excludes elements (from start until end) from a slice
arguments:
	slice: the slice to exclude from
	start: the starting index of the exclusion (inclusive)
	end: the ending index of the exclusion (exclusive)
retunr: the excluded from slice
*/
func ExcludeSliceElements(slice []any, start, end uint) []any {
	var result []any

	result = append(result, slice[:start + 1]...)
	result = append(result, slice[end:]...)

	return result
}