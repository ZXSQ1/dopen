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
