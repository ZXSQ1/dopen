package argp

type OptionHandler func(...string)

type ArgParser struct {
	args            []string                 // program arguments
	optionHandlers  map[string]OptionHandler // option handlers
	optionArgLength map[string]uint          // a map of options and the number of their arguments
}

/*
description: gets a new ArgParser object
arguments:

	args: a string slice of the program arguments

return: an ArgParser object
*/
func GetArgParser(args []string) ArgParser {
	return ArgParser{
		args: args,
	}
}

/*
description: assigns option handlers to specified arguments
arguments:

	options: a string slice containing the options to handle
	fn: the option handler function of the OptionHandler type
	argLength: the number of arguments that go after a specified option

return:
*/
func (argParser *ArgParser) HandleArgs(options []string, fn OptionHandler, argLength uint) {
	for _, option := range options {
		argParser.optionHandlers[option] = fn
		argParser.optionArgLength[option] = argLength
	}
}
