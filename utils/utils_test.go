package utils

import (
	"os/exec"
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
