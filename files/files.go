package files

import (
	"fmt"
	"io"
	"os"
)

// Errors

var ErrNotExist error = fmt.Errorf("file not exists")
var ErrNotFile error = fmt.Errorf("path not file")
var ErrNotDir error = fmt.Errorf("path not dir")

var ErrOpenFile error = fmt.Errorf("error opening file")
var ErrCreateFile error = fmt.Errorf("error creating file")
var ErrWriteFile error = fmt.Errorf("error writing file")
var ErrReadFile error = fmt.Errorf("error reading file")

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
		return false, ErrNotExist
	}

	stat, _ := os.Stat(file)
	return stat.IsDir(), nil
}

/*
description: checks if the file is a file
arguments:

	file: the file path to check for

return:
  - true if it is a file and false otherwise
  - an error if it doesn't exist
*/
func IsFile(file string) (bool, error) {
	isDir, err := IsDir(file)

	if err != nil {
		return !isDir, nil
	} else {
		return false, err
	}
}

/*
description: used to get the file object
arguments:

	file: the file path to return an object from

return:
	- the file object
	- an error
*/
func GetFile(file string) (result *os.File, retErr error) {
	isDir, _ := IsDir(file)
	isFile := !isDir

	if IsExists(file) && isFile {
		fileObj, err := os.Open(file)

		if err != nil {
			return nil, ErrOpenFile
		}

		result = fileObj
	} else if !IsExists(file) {
		fileObj, err := os.Create(file)

		if err != nil {
			return nil, ErrCreateFile
		}

		result = fileObj
	} else if IsExists(file) && isDir {
		return nil, ErrNotFile
	}

	return
}

/*
description: writes data to a file
arguments:

	file: the file path to write to
	data: the data to write to the file

return: an error if anything goes wrong
*/
func WriteFile(file string, data []byte) error {
	fileObj, err := GetFile(file)
	defer fileObj.Close()

	if err != nil {
		return err
	}

	_, err = fileObj.Write(data)

	if err != nil {
		return ErrWriteFile
	}

	return nil
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
