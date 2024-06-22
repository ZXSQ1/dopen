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

func TestArgParser_Execute(t *testing.T) {
	input := []string{"kitty", "-d", "-b", "--hold", "bash", "-c", "source ~/.bashrc; clear"}
	argParser := GetArgParser(input)

	argParser.HandleArgs([]string{"-b", "-d"}, func(s ...string) {
		println("test 1 successful")
	}, 0)

	argParser.HandleArgs([]string{"--hold"}, func(s ...string) {
		println("test 2 successful")
	}, 3)

	args := argParser.Execute()

	if !slices.Equal(args, []string{"kitty"}) {
		for _, value := range args {
			println(value)
		}
		
		t.Fail()
	}
}