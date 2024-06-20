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

func TestArgParser_HandleArgs(t *testing.T) {
	input := []string{"cp", "-p", "-r", "-P", "name", "type", "file/path"}
	argParser := GetArgParser(input)

	argParser.HandleArgs([]string{"-p", "-r"}, func(s ...string) {}, 0)
	argParser.HandleArgs([]string{"-P"}, func(s ...string) {}, 2)

	for _, val := range []string{"-p", "-r", "-P"} {
		if _, found := argParser.optionHandlers[val]; !found {
			t.Fail()
		}
	}

	if argParser.optionArgLength["-p"] != 0 || argParser.optionArgLength["-P"] != 2 {
		t.Fail()
	}

}
