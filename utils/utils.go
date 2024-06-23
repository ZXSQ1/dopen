package utils

import (
	"io"
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
	Message  []byte
	Position int
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
	messenger.Message = append(messenger.Message, p...)
	return len(p), nil
}

/*
description: reads from the messenger into the buffer
arguments:

	p: the byte to slice to read into

return:
  - the number of bytes read
  - an error object
*/
func (messenger *Messenger) Read(p []byte) (n int, err error) {
	var length = 0
	var messageLength = len(messenger.Message[messenger.Position:])

	if messageLength == 0 {
		return 0, io.EOF
	} else if len(p) < messageLength {
		length = len(p)
	} else if len(p) > messageLength {
		length = messageLength
	} else {
		length = len(p)
	}

	for i := 0; i < length; i++ {
		p[i] = messenger.Message[messenger.Position + i]
	}

	messenger.Position += length

	return length, nil
}
