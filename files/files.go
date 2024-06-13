package files

import (
	"log"
	"os"
)

/*
description: checks if the file exists or not
arguments:

	file: the file path to check existence

return: true if it exists and false otherwise
*/
func IsExists(file string) bool {
  _, err := os.Stat(file)

	return !os.IsNotExist(err)
}

/*
description: checks if the file is a directory
arguments:

  file: the file path to check for existence

return: true if it is a directory and false otherwise
*/
func IsDir(file string) bool {
  if !IsExists(file) {
    log.Fatalln("IsDir: file not exists")
  }

  stat, _ := os.Stat(file)
  return stat.IsDir()
}
