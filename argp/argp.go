package argp

type ArgParser struct {
	args           []string // arguments except arg[0]
	optionHandlers map[string]func(...[]string) // option handlers
}
