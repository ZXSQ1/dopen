package files

import (
	"fmt"
	"io"
	"os"
)

// Errors

var errNotExist error = fmt.Errorf("file not exists")
var errNotFile error = fmt.Errorf("path not file")
var errNotDir error = fmt.Errorf("path not dir")

////////

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

return:
  - true if it is a directory and false otherwise
  - an error if it doesn't exist
*/
func IsDir(file string) (bool, error) {
	if !IsExists(file) {
		return false, errNotExist
	}

	stat, _ := os.Stat(file)
	return stat.IsDir(), nil
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

/*
description: used to get the file object
arguments:

	file: the file path to return an object from

return: the file object
*/
func GetFile(file string) (result *os.File) {
	if IsExists(file) && IsFile(file) {
		fileObj, err := os.Open(file)

		if err != nil {
			fmt.Printf("GetFile: error opening file %s\n", file)
		}

		result = fileObj
	} else if !IsExists(file) {
		fileObj, err := os.Create(file)

		if err != nil {
			fmt.Printf("GetFile: error creating file %s\n", file)
		}

		result = fileObj
	} else if IsExists(file) && IsDir(file) {
		fmt.Printf("GetFile: file %s is a directory\n", file)
	}

	return
}

/*
description: writes data to a file
arguments:

	file: the file path to write to
	data: the data to write to the file

return:
*/
func WriteFile(file string, data []byte) {
	fileObj := GetFile(file)
	defer fileObj.Close()

	_, err := fileObj.Write(data)

	if err != nil {
		fmt.Printf("WriteFile: error writing file %s\n", file)
	}
}

/*
description: reads data from a file
arguments:

	file: the file to read from

return: the read bytes
*/
func ReadFile(file string) (result []byte) {
	fileObj := GetFile(file)
	defer fileObj.Close()

	buffer := make([]byte, 1024)

	for {
		nBytes, err := fileObj.Read(buffer)
		result = append(result, buffer[:nBytes]...)

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("ReadFile: %s\n", err.Error())
		}
	}

	return
}
