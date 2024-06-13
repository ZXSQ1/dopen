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

  file: the file path to check for

return: true if it is a directory and false otherwise
*/
func IsDir(file string) bool {
  if !IsExists(file) {
    log.Fatalln("IsDir: file not exists")
  }

  stat, _ := os.Stat(file)
  return stat.IsDir()
}

/*
description: checks if the file is a file
arguments:

  file: the file path to check for

return: true if it is a file and false otherwise
*/
func IsFile(file string) bool {
  return !IsDir(file)
}

func WriteFile(file string, data []byte) {
  var fileObj *os.File

  if IsExists(file) && IsFile(file) {
    fileObj, err := os.Open(file)
    
    if err != nil {
      log.Fatalf("WriteFile: error opening file %s\n", file)
    }

    defer fileObj.Close()
  } else if !IsExists(file) {
    fileObj, err := os.Create(file)

    if err != nil {
      log.Fatalf("WriteFile: error creating file %s\n", file)
    }

    defer fileObj.Close()
  } else if IsExists(file) && IsDir(file) {
    log.Fatalf("WriteFile: file %s is a directory\n", file)
  } 

  _, err := fileObj.Write(data);

  if err != nil {
    log.Fatalf("WriteFile: error writing file %s\n", file)
  }
}
