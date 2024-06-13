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

func IsDir(file string) bool {
  if !IsExists(file) {
    log.Fatalln("IsDir: file not exists")
  }

  stat, _ := os.Stat(file)
  return stat.IsDir()
}
