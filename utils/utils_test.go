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