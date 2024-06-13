package files

import (
	"os"
)

/*
description: checks if the file exists or not
arguments:

	file: the file path to check existence

return: true if it exists and false otherwise
*/
func IsExists(file string) bool {
	return !os.IsNotExist(file)
}

func IsFile(file string) bool {
	if !IsExists(file) {
		return false
	}

	stat, _ := os.Stat(file)

  return !stat.IsDir()
}
