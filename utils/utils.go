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

type Messenger struct {
	message []byte
}

/*
description: writes to the messenger
arguments:

	p: the byte slice to write

return:
  - the number of bytes written
  - an error object
*/
func (messenger *Messenger) Write(p []byte) (n int, err error) {
	messenger.message = p
	return len(p), nil
}
