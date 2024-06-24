package main

import (
	"os"
	"strings"
	"unicode"

	"github.com/ZXSQ1/dopen/argp"
	"github.com/ZXSQ1/dopen/doc_manager"
)

func help() {
	message := `
		syntax: <program> <language-doc-name>
		options:
			--help, -h 			for help
			--width, -w	<uint>	for setting column width
	`

	message = strings.ReplaceAll(message, "\t", " ")
	message = strings.Join(strings.Split(message, "\n")[1:], "\n")

	println(message)
	os.Exit(1)
}

func start(args []string) {
	if len(args) < 2 {
		help()
	} else {
		args = args[1:]
		argParser := argp.GetArgParser(args)

		argParser.HandleArgs([]string{"-h", "--help"}, func(s ...string) { help() }, 0)
		argParser.HandleArgs([]string{"-w", "--width"}, func(s ...string) {
			if len(s) < 1 {
				help()
			}

			for _, char := range s[0] {
				if !unicode.IsDigit(char) {
					help()
				}
			}

			doc_manager.ColumnWidth = s[0]
		}, 1)

		args = argParser.Execute()

		if len(args) < 1 {
			help()
		}

		for _, arg := range args {
			doc_manager.OpenDocs(arg)
		}
	}
}

func main() {
	argv := os.Args

	start(argv)
}
