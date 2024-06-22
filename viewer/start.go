package viewer

import (
	"os"

	"github.com/ZXSQ1/devdocs-tui/argp"
)

const (
	optionINSTALL = 0
	optionHELP    = 1
	optionLIST    = 2
	optionSEARCH  = 3
	optionFETCH   = 4
	optionREMOVE  = 5
)

/*
description: used in routing and handling options
arguments:
return:
*/
func Route() {
	argv := os.Args

	if len(argv) < 2 {
		// help
	} else {
		argParser := argp.GetArgParser(argv)
		option := optionHELP
		optionPtr := &option

		argParser.HandleArgs([]string{"-i", "--install"}, func(s ...string) { *optionPtr = optionINSTALL }, 0)
		argParser.HandleArgs([]string{"-s", "--search"}, func(s ...string) { *optionPtr = optionSEARCH }, 0)
		argParser.HandleArgs([]string{"-l", "--list"}, func(s ...string) { *optionPtr = optionLIST }, 0)
		argParser.HandleArgs([]string{"-h", "--help"}, func(s ...string) { *optionPtr = optionHELP }, 0)
		argParser.HandleArgs([]string{"-f", "--fetch"}, func(s ...string) { *optionPtr = optionFETCH }, 0)
		argParser.HandleArgs([]string{"-r", "--remove"}, func(s ...string) { *optionPtr = optionREMOVE }, 0)

		argParser.Execute()

		switch option {
		case optionHELP:
			// help
		case optionINSTALL:
			// install
		case optionLIST:
			// list
		case optionSEARCH:
			// search
		case optionFETCH:
			// fetch
		case optionREMOVE:
			// remove
		default:
			// help
		}
	}
}
