package argp

import (
	"slices"
	"testing"
)

func TestGetArgParser(t *testing.T) {
	input := []string{"hi", "what", "is"}
	argParser := GetArgParser(input)

	if !slices.Equal(argParser.args, input) {
		t.Fail()
	}
}

