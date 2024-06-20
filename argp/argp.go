package argp

type ArgParser struct {
	args           []string // program arguments
	optionHandlers map[string]func(...[]string) // option handlers
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