package utils

import (
	"os/exec"
	"slices"
	"strings"
	"testing"
)

func TestGetEnvironVar(t *testing.T) {
	variableName := "HOME"

	out, _ := exec.Command("bash", "-c", "echo $"+variableName).Output()
	expectedResult := strings.TrimSpace(string(out))
	actualResult := GetEnvironVar(variableName)

	println(expectedResult, actualResult)

	if expectedResult != actualResult {
		t.Fail()
	}
}

func TestMessenger_Write(t *testing.T) {
	message := []byte("hello")
	messenger := Messenger{}

	n, err := messenger.Write(message)

	if err != nil || n != len(message) || !slices.Equal(messenger.Message, message) {
		t.Fail()
	}
}

func TestMessenger_Read(t *testing.T) {
	message := []byte("Hallo! Wie heiSSt du? Ich heiSSe *name*. This has been German.")
	messenger := Messenger{}

	messenger.Write(message)
	result := make([]byte, len(message))
	buffer := make([]byte, 5)

skip:
	for {
		n, _ := messenger.Read(buffer)

		if n == 0 {
			break skip
		}

		result = append(result, buffer[:n]...)
	}

	print(string(result))

	if !slices.Equal(result, message) {
		t.Fail()
	}
}
